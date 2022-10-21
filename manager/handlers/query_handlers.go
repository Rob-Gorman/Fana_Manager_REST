package handlers

import (
	"fmt"
	"manager/models"
	"manager/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (h Handler) GetAllFlags(w http.ResponseWriter, r *http.Request) {
	var flags []models.Flag

	result := h.DB.Find(&flags)

	if result.Error != nil {
		fmt.Println(result.Error)
	}

	response := []models.FlagNoAudsResponse{}

	for ind := range flags {
		response = append(response, models.FlagNoAudsResponse{Flag: &flags[ind]})
	}

	utils.PayloadResponse(w, r, &response)

}

func (h Handler) GetAllAudiences(w http.ResponseWriter, r *http.Request) {
	auds := []models.Audience{}
	respAuds := []models.AudienceNoCondsResponse{}

	result := h.DB.Preload("Conditions").Find(&auds)

	if result.Error != nil {
		fmt.Println(result.Error)
	}

	for ind := range auds {
		respAuds = append(respAuds, models.AudienceNoCondsResponse{Audience: &auds[ind]})
	}

	utils.PayloadResponse(w, r, respAuds)
}

func (h Handler) GetAllAttributes(w http.ResponseWriter, r *http.Request) {
	var attrs []models.Attribute

	result := h.DB.Find(&attrs)

	if result.Error != nil {
		fmt.Println(result.Error)
	}

	utils.PayloadResponse(w, r, attrs)
}

func (h Handler) GetFlag(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Invalid flag ID."))
		return
	}

	var flag models.Flag
	auds := []models.AudienceNoCondsResponse{}

	err = h.DB.Preload("Audiences").First(&flag, id).Error
	if err != nil {
		utils.NoRecordResponse(w, r, err)
		return
	}

	for ind := range flag.Audiences {
		auds = append(auds, models.AudienceNoCondsResponse{Audience: &flag.Audiences[ind]})
	}

	utils.PayloadResponse(w, r, &models.FlagResponse{Flag: &flag, Audiences: auds})
}

func (h Handler) GetAudience(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Invalid audience ID."))
		return
	}

	var aud models.Audience

	err = h.DB.Preload("Flags").Preload("Conditions").First(&aud, id).Error

	if err != nil {
		utils.NoRecordResponse(w, r, err)
		return
	}

	conds := GetEmbeddedConds(aud, h.DB)
	flags := GetEmbeddedFlags(aud.Flags)

	response := models.AudienceResponse{
		Audience:   &aud,
		Conditions: conds,
		Flags:      flags,
	}

	utils.PayloadResponse(w, r, &response)
}

func (h Handler) GetAttribute(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Invalid attribute ID."))
		return
	}

	var attr models.Attribute

	err = h.DB.Preload("Conditions").First(&attr, id).Error

	if err != nil {
		utils.NoRecordResponse(w, r, err)
		return
	}

	response := BuildAttrResponse(attr, h)

	utils.PayloadResponse(w, r, &response)
}

func (h Handler) GetAuditLogs(w http.ResponseWriter, r *http.Request) {
	flags := []models.FlagLog{}
	h.DB.Find(&flags)

	auds := []models.AudienceLog{}
	h.DB.Find(&auds)

	attrs := []models.AttributeLog{}
	h.DB.Find(&attrs)

	response := models.AuditResponse{
		FlagLogs:      flags,
		AudienceLogs:  auds,
		AttributeLogs: attrs,
	}

	utils.PayloadResponse(w, r, &response)
}

func (h Handler) GetSdkKeys(w http.ResponseWriter, r *http.Request) {
	sdks := []models.Sdkkey{}
	h.DB.Find(&sdks)
	utils.PayloadResponse(w, r, &sdks)
}
