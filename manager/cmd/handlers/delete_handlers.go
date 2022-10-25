package handlers

import (
	"manager/utils"
	"net/http"
)

func (h Handler) DeleteFlag(w http.ResponseWriter, r *http.Request) {
	id, err := h.idFromParams(w, r, "flag")
	if err != nil {
		return
	}

	code, err := h.DM.DeleteFlag(id)
	if err != nil {
		utils.ErrorResponse(w, r, code, err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h Handler) DeleteAudience(w http.ResponseWriter, r *http.Request) {
	id, err := h.idFromParams(w, r, "audience")
	if err != nil {
		return
	}

	code, err := h.DM.DeleteAudience(id)
	if err != nil {
		utils.ErrorResponse(w, r, code, err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h Handler) DeleteAttribute(w http.ResponseWriter, r *http.Request) {
	id, err := h.idFromParams(w, r, "attribute")
	if err != nil {
		return
	}

	code, err := h.DM.DeleteAttribute(id)
	if err != nil {
		utils.ErrorResponse(w, r, code, err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h Handler) RegenSDKkey(w http.ResponseWriter, r *http.Request) {
	id, err := h.idFromParams(w, r, "sdk key")
	if err != nil {
		return
	}

	res, code, err := h.DM.RegenSDKkey(id)
	if err != nil {
		utils.ErrorResponse(w, r, code, err.Error())
		return
	}

	h.ProcessServices(nil, "")
	h.ComposeResponse(w, r, res, err)
}
