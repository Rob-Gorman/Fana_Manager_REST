package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"manager/cache"
	"manager/models"
	"manager/publisher"
	"manager/utils"
	"math/rand"
	"time"

	"gorm.io/gorm"
)

func GetEmbeddedConds(aud models.Audience, db *gorm.DB) []models.ConditionEmbedded {
	conds := []models.ConditionEmbedded{}
	for ind := range aud.Conditions {
		cond := aud.Conditions[ind]
		var attr models.Attribute
		db.Find(&attr, cond.AttributeID)
		db.Find(&cond)
		cond.Attribute = attr
		conds = append(conds, models.ConditionEmbedded{
			Condition:    &cond,
			AttributeKey: attr.Key,
		})
	}
	return conds
}

func FlagReqToFlag(flagReq models.FlagSubmit, h Handler) (flag models.Flag) {
	auds := []models.Audience{}

	h.DB.Where("key in (?)", flagReq.Audiences).Find(&auds)

	flag = models.Flag{
		Audiences:   auds,
		Key:         flagReq.Key,
		DisplayName: flagReq.DisplayName,
		Sdkkey:      flagReq.SdkKey,
	}

	return flag
}

func FlagToFlagResponse(flag models.Flag, h Handler) models.FlagResponse {
	h.DB.Preload("Audiences").First(&flag)
	respAuds := []models.AudienceNoCondsResponse{}
	for ind := range flag.Audiences {
		respAuds = append(respAuds, models.AudienceNoCondsResponse{Audience: &flag.Audiences[ind]})
	}
	return models.FlagResponse{
		Flag:      &flag,
		Audiences: respAuds,
	}
}

func BuildAudUpdate(req models.Audience, id int, h Handler) (aud models.Audience) {
	h.DB.Find(&aud, id)
	aud.Conditions = req.Conditions
	aud.Combine = req.Combine
	aud.DisplayName = req.DisplayName
	return aud
}

func GetEmbeddedFlags(flags []models.Flag) []models.FlagNoAudsResponse {
	fr := []models.FlagNoAudsResponse{}
	for i := range flags {
		fr = append(fr, models.FlagNoAudsResponse{Flag: &flags[i]})
	}

	return fr
}

func BuildAttrResponse(a models.Attribute, h Handler) models.AttributeResponse {
	conds := a.Conditions
	audids := []uint{}
	for _, cond := range conds {
		audids = append(audids, cond.AudienceID)
	}

	auds := []models.Audience{}

	if len(audids) > 0 {
		h.DB.Find(&auds, audids)
	}

	respauds := []models.AudienceNoCondsResponse{}
	for i := range auds {
		respauds = append(respauds, models.AudienceNoCondsResponse{
			Audience: &auds[i],
		})
	}

	return models.AttributeResponse{
		Attribute: &a,
		Audiences: respauds,
	}
}

func PublishContent(data interface{}, channel string) {
	if publisher.Redis == nil {
		utils.HandleErr(nil, "No connection established with Redis; cannot publish")
		return
	}

	byteArray, err := json.Marshal(data)
	if err != nil {
		utils.HandleErr(err, "Unmarshalling error")
	}
	fmt.Println("Manager trying to publish", string(byteArray))

	err = publisher.Redis.Publish(context.TODO(), channel, byteArray).Err()
	if err != nil {
		utils.HandleErr(err, " : Error trying to publish to redis")
		return
	}
}

func RefreshCache(db *gorm.DB) {
	flagCache := cache.InitFlagCache()
	fs := BuildFlagset(db)
	flagCache.FlushAllAsync()
	flagCache.Set("data", &fs)
}

func OrphanedAud(aud *models.Audience) bool {
	return len((*aud).Flags) == 0
}

func OrphanedAttr(attr *models.Attribute, h Handler) bool {
	asscs := h.DB.Model(attr).Association("Conditions").Count()
	return asscs == 0
}

func NewSDKKey(s string) string {
	digits := []byte("0123456789abcdefghijkm")
	newKey := []byte{}
	rand.Seed(time.Now().UnixNano())

	for _, char := range s {
		if char == '-' {
			newKey = append(newKey, '-')
		} else {
			randInd := rand.Intn(len(digits))
			newKey = append(newKey, digits[randInd])
		}
	}

	return string(newKey)
}
