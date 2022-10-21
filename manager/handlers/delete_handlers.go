package handlers

import (
	"manager/models"
	"manager/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (h Handler) DeleteFlag(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Invalid flag ID."))
		return
	}

	flag := &models.Flag{}
	err = h.DB.Preload("Audiences").First(&flag, id).Error
	if err != nil {
		utils.NoRecordResponse(w, r, err)
		return
	}

	h.DB.Model(&flag).Association("Audiences").Delete(flag.Audiences)
	err = h.DB.Unscoped().Delete(&flag).Error
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h Handler) DeleteAudience(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Invalid audience ID."))
		return
	}

	aud := &models.Audience{}
	err = h.DB.Preload("Flags").First(&aud, id).Error
	if err != nil {
		utils.NoRecordResponse(w, r, err)
		return
	}
	if !OrphanedAud(aud) {
		msg := "Cannot delete Audience assigned to Flags."
		utils.UnprocessableEntityResponse(w, r, nil, msg)
		return
	}

	h.DB.Model(&aud).Association("Flags").Delete(aud.Flags)
	err = h.DB.Unscoped().Delete(&aud).Error
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h Handler) DeleteAttribute(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Invalid attribute ID."))
		return
	}

	attr := &models.Attribute{}
	err = h.DB.First(attr, id).Error
	if err != nil {
		utils.NoRecordResponse(w, r, err)
		return
	}
	if !OrphanedAttr(attr, h) {
		msg := "Cannot delete Attribute assigned to Audiences."
		utils.UnprocessableEntityResponse(w, r, nil, msg)
		return
	}

	err = h.DB.Unscoped().Delete(&attr).Error
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h Handler) RegenSDKkey(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Invalid flag ID."))
		return
	}

	sdk := models.Sdkkey{}
	h.DB.Find(&sdk, id)

	newSDK := models.Sdkkey{
		Key:  NewSDKKey(sdk.Key),
		Type: sdk.Type,
	}

	err = h.DB.Create(&newSDK).Error
	if err != nil {
		utils.UnavailableResponse(w, r, err)
		return
	}

	h.DB.Unscoped().Delete(&sdk)

	h.DB.Find(&newSDK)

	RefreshCache(h.DB)

	utils.CreatedResponse(w, r, &newSDK)
}
