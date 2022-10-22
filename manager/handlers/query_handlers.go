package handlers

import (
	"manager/models"
	"manager/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (h Handler) GetAllFlags(w http.ResponseWriter, r *http.Request) {
	res, err := h.DM.GetAllFlags()

	h.ComposeResponse(w, r, res, err)
}

func (h Handler) GetAllAudiences(w http.ResponseWriter, r *http.Request) {
	res, err := h.DM.GetAllAudiences()

	h.ComposeResponse(w, r, res, err)
}

func (h Handler) GetAllAttributes(w http.ResponseWriter, r *http.Request) {
	res, err := h.DM.GetAllAttributes()

	h.ComposeResponse(w, r, res, err)
}

func (h Handler) GetFlag(w http.ResponseWriter, r *http.Request) {
	idParam := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		utils.MalformedIDResponse(w, r, "flag", idParam)
		return
	}

	res, err := h.DM.GetFlag(id)

	h.ComposeResponse(w, r, res, err)
}

// TODO
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

	res := models.AudienceResponse{
		Audience:   &aud,
		Conditions: conds,
		Flags:      flags,
	}

	utils.PayloadResponse(w, r, &res)
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

	res := BuildAttrResponse(attr, h)

	utils.PayloadResponse(w, r, &res)
}

func (h Handler) GetAuditLogs(w http.ResponseWriter, r *http.Request) {
	flags := []models.FlagLog{}
	h.DB.Find(&flags)

	auds := []models.AudienceLog{}
	h.DB.Find(&auds)

	attrs := []models.AttributeLog{}
	h.DB.Find(&attrs)

	res := models.AuditResponse{
		FlagLogs:      flags,
		AudienceLogs:  auds,
		AttributeLogs: attrs,
	}

	utils.PayloadResponse(w, r, &res)
}

func (h Handler) GetSdkKeys(w http.ResponseWriter, r *http.Request) {
	sdks := []models.Sdkkey{}
	h.DB.Find(&sdks)
	utils.PayloadResponse(w, r, &sdks)
}
