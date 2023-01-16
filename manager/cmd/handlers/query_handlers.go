package handlers

import (
	"manager/utils"
	"net/http"
)

func (h Handler) GetAllFlags(w http.ResponseWriter, r *http.Request) {
	res, err := h.DM.GetAllFlags()

	if err != nil {
		utils.ErrLog.Falalf("failed to return flags %v", err)
	}

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
	id, err := h.idFromParams(w, r, "flag")
	if err != nil {
		return
	}

	res, err := h.DM.GetFlag(id)

	h.ComposeResponse(w, r, res, err)
}

func (h Handler) GetAudience(w http.ResponseWriter, r *http.Request) {
	id, err := h.idFromParams(w, r, "audience")
	if err != nil {
		return
	}

	res, err := h.DM.GetAudience(id)

	h.ComposeResponse(w, r, res, err)
}

func (h Handler) GetAttribute(w http.ResponseWriter, r *http.Request) {
	id, err := h.idFromParams(w, r, "attribute")
	if err != nil {
		return
	}

	res, err := h.DM.GetAttribute(id)
	h.ComposeResponse(w, r, res, err)
}

func (h Handler) GetAuditLogs(w http.ResponseWriter, r *http.Request) {
	res, err := h.DM.GetAuditLogs()

	h.ComposeResponse(w, r, res, err)
}

func (h Handler) GetSdkKeys(w http.ResponseWriter, r *http.Request) {
	res, err := h.DM.GetSdkKeys()

	h.ComposeResponse(w, r, res, err)
}
