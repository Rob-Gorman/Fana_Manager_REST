package datamodel

import (
	"manager/internal/cache"
	"manager/internal/data/models"
	"math/rand"
	"time"
)

func (d *DataModel) GetEmbeddedConds(aud models.Audience) []models.ConditionEmbedded {
	db := d.DB

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

func (d *DataModel) GetEmbeddedFlags(flags []models.Flag) []models.FlagNoAudsResponse {
	fr := []models.FlagNoAudsResponse{}
	for i := range flags {
		fr = append(fr, models.FlagNoAudsResponse{Flag: &flags[i]})
	}

	return fr
}

// refactor out DB dependency?
func (d *DataModel) FlagReqToFlag(flagReq models.FlagSubmit) (flag models.Flag) {
	auds := []models.Audience{}

	d.DB.Where("key in (?)", flagReq.Audiences).Find(&auds)

	flag = models.Flag{
		Audiences:   auds,
		Key:         flagReq.Key,
		DisplayName: flagReq.DisplayName,
		Sdkkey:      flagReq.SdkKey,
	}

	return flag
}

func (d *DataModel) FlagToFlagResponse(flag models.Flag) models.FlagResponse {
	d.DB.Preload("Audiences").First(&flag)
	respAuds := []models.AudienceNoCondsResponse{}
	for ind := range flag.Audiences {
		respAuds = append(respAuds, models.AudienceNoCondsResponse{Audience: &flag.Audiences[ind]})
	}
	return models.FlagResponse{
		Flag:      &flag,
		Audiences: respAuds,
	}
}

func buildAudUpdate(req models.Audience, id int, d *DataModel) (aud models.Audience) {
	d.DB.Find(&aud, id)
	aud.Conditions = req.Conditions
	aud.Combine = req.Combine
	aud.DisplayName = req.DisplayName
	return aud
}

func (d *DataModel) BuildAttrResponse(a models.Attribute) models.AttributeResponse {
	conds := a.Conditions
	audids := []uint{}
	for _, cond := range conds {
		audids = append(audids, cond.AudienceID)
	}

	auds := []models.Audience{}

	if len(audids) > 0 {
		d.DB.Find(&auds, audids)
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

// fix this function / prob extract
func (d *DataModel) RefreshCache() {
	flagCache := cache.InitFlagCache()
	fs := d.BuildFlagset()
	flagCache.FlushAllAsync()
	flagCache.Set("data", &fs)
}

func OrphanedAud(aud *models.Audience) bool {
	return len(aud.Flags) == 0
}

func OrphanedAttr(attr *models.Attribute) bool {
	return len(attr.Conditions) == 0
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
