package datamodel

import (
	"manager/data/models"
	"manager/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (d *DataModel) DeleteFlag(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Invalid flag ID."))
		return
	}

	flag := &models.Flag{}
	err = d.DB.Preload("Audiences").First(&flag, id).Error
	if err != nil {
		utils.NoRecordResponse(w, r, err)
		return
	}

	d.DB.Model(&flag).Association("Audiences").Delete(flag.Audiences)
	err = d.DB.Unscoped().Delete(&flag).Error
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (d *DataModel) DeleteAudience(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Invalid audience ID."))
		return
	}

	aud := &models.Audience{}
	err = d.DB.Preload("Flags").First(&aud, id).Error
	if err != nil {
		utils.NoRecordResponse(w, r, err)
		return
	}
	if !OrphanedAud(aud) {
		msg := "Cannot delete Audience assigned to Flags."
		utils.UnprocessableEntityResponse(w, r, nil, msg)
		return
	}

	d.DB.Model(&aud).Association("Flags").Delete(aud.Flags)
	err = d.DB.Unscoped().Delete(&aud).Error
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (d *DataModel) DeleteAttribute(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Invalid attribute ID."))
		return
	}

	attr := &models.Attribute{}
	err = d.DB.First(attr, id).Error
	if err != nil {
		utils.NoRecordResponse(w, r, err)
		return
	}
	if !OrphanedAttr(attr, d) {
		msg := "Cannot delete Attribute assigned to Audiences."
		utils.UnprocessableEntityResponse(w, r, nil, msg)
		return
	}

	err = d.DB.Unscoped().Delete(&attr).Error
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (d *DataModel) RegenSDKkey(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Invalid flag ID."))
		return
	}

	sdk := models.Sdkkey{}
	d.DB.Find(&sdk, id)

	newSDK := models.Sdkkey{
		Key:  NewSDKKey(sdk.Key),
		Type: sdk.Type,
	}

	err = d.DB.Create(&newSDK).Error
	if err != nil {
		utils.UnavailableResponse(w, r, err)
		return
	}

	d.DB.Unscoped().Delete(&sdk)

	d.DB.Find(&newSDK)

	RefreshCache(d.DB)

	utils.CreatedResponse(w, r, &newSDK)
}
