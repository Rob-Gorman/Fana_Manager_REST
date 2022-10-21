package handlers

import (
	"encoding/json"
	"io/ioutil"
	"manager/models"
	"manager/utils"
	"net/http"

	"gorm.io/gorm"
)

func (h Handler) CreateFlag(w http.ResponseWriter, r *http.Request) {
	var flagReq models.FlagSubmit

	// defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)

	err := json.Unmarshal(body, &flagReq)
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	flag := FlagReqToFlag(flagReq, h)

	err = h.DB.Omit("Audiences.*").Session(&gorm.Session{FullSaveAssociations: true}).Create(&flag).Error

	if err != nil {
		msg := "Cannot create resource with duplicate keys."
		utils.UnprocessableEntityResponse(w, r, err, msg)
		return
	}

	h.DB.Preload("Audiences").Find(&flag)
	respAuds := []models.AudienceNoCondsResponse{}

	for ind := range flag.Audiences {
		respAuds = append(respAuds, models.AudienceNoCondsResponse{
			Audience: &flag.Audiences[ind],
		})
	}

	response := models.FlagResponse{
		Flag:      &flag,
		Audiences: respAuds,
	}
	PublishContent(&response, "flag-update-channel")

	utils.CreatedResponse(w, r, &response)
	RefreshCache(h.DB)
}

func (h Handler) CreateAttribute(w http.ResponseWriter, r *http.Request) {
	var attrReq models.Attribute
	// defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	err = json.Unmarshal(body, &attrReq)
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	err = h.DB.Create(&attrReq).Error
	if err != nil {
		msg := "Cannot create resource with duplicate keys."
		utils.UnprocessableEntityResponse(w, r, err, msg)
		return
	}

	h.DB.Find(&attrReq)
	utils.CreatedResponse(w, r, &attrReq)
	RefreshCache(h.DB)
}

func (h Handler) CreateAudience(w http.ResponseWriter, r *http.Request) {
	var aud models.Audience

	// defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	err = json.Unmarshal(body, &aud)
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	err = h.DB.Session(&gorm.Session{FullSaveAssociations: true}).Create(&aud).Error
	if err != nil {
		msg := "Cannot create resource with duplicate keys."
		utils.UnprocessableEntityResponse(w, r, err, msg)
		return
	}

	h.DB.Model(&models.Audience{}).Preload("Conditions").Find(&aud)

	response := models.AudienceResponse{
		Audience:   &aud,
		Conditions: GetEmbeddedConds(aud, h.DB),
		Flags:      []models.FlagNoAudsResponse{},
	}

	PublishContent(&aud, "audience-update-channel")
	RefreshCache(h.DB)

	utils.CreatedResponse(w, r, &response)
}
