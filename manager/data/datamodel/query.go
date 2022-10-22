package datamodel

import (
	"fmt"
	"manager/data/models"
	"manager/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (d *DataModel) GetAllFlags() (*[]byte, error) {
	var flags *[]models.Flag
	err := d.DB.Find(flags).Error

	if err != nil {
		return nil, err
	}

	res := models.AllFlagsRes(flags)
	return models.ToJSON(res)
}

func (d *DataModel) GetAllAudiences() (*[]byte, error) {
	auds := []models.Audience{}

	err := d.DB.Preload("Conditions").Find(&auds).Error
	if err != nil {
		return nil, err
	}

	res := models.AllAudsRes(&auds)
	return models.ToJSON(res)
}

func (d *DataModel) GetAllAttributes() (*[]byte, error) {
	var attrs []models.Attribute

	err := d.DB.Find(&attrs).Error
	if err != nil {
		return nil, err
	}

	return models.ToJSON(attrs)
}

func (d *DataModel) GetFlag(id int) (*[]byte, error) {
	var flag models.Flag
	auds := []models.AudienceNoCondsResponse{}

	err := d.DB.Preload("Audiences").First(&flag, id).Error
	if err != nil {
		return nil, err
	}

	for ind := range flag.Audiences {
		auds = append(auds, models.AudienceNoCondsResponse{Audience: &flag.Audiences[ind]})
	}

	res := &models.FlagResponse{Flag: &flag, Audiences: auds}
	return models.ToJSON(res)
}

func (d *DataModel) GetAudience() (*[]byte, error) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Invalid audience ID."))
		return
	}

	var aud models.Audience

	err = d.DB.Preload("Flags").Preload("Conditions").First(&aud, id).Error

	if err != nil {
		utils.NoRecordResponse(w, r, err)
		return
	}

	conds := GetEmbeddedConds(aud, d.DB)
	flags := GetEmbeddedFlags(aud.Flags)

	response := models.AudienceResponse{
		Audience:   &aud,
		Conditions: conds,
		Flags:      flags,
	}

	utils.PayloadResponse(w, r, &response)
}

func (d *DataModel) GetAttribute() (*[]byte, error) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Invalid attribute ID."))
		return
	}

	var attr models.Attribute

	err = d.DB.Preload("Conditions").First(&attr, id).Error

	if err != nil {
		utils.NoRecordResponse(w, r, err)
		return
	}

	response := BuildAttrResponse(attr, d)

	utils.PayloadResponse(w, r, &response)
}

func (d *DataModel) GetAuditLogs() (*[]byte, error) {
	flags := []models.FlagLog{}
	d.DB.Find(&flags)

	auds := []models.AudienceLog{}
	d.DB.Find(&auds)

	attrs := []models.AttributeLog{}
	d.DB.Find(&attrs)

	response := models.AuditResponse{
		FlagLogs:      flags,
		AudienceLogs:  auds,
		AttributeLogs: attrs,
	}

	utils.PayloadResponse(w, r, &response)
}

func (d *DataModel) GetSdkKeys(w http.ResponseWriter, r *http.Request) {
	sdks := []models.Sdkkey{}
	d.DB.Find(&sdks)
	utils.PayloadResponse(w, r, &sdks)
}
