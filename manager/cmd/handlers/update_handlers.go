package handlers

import (
	"fmt"
	"io/ioutil"
	"manager/utils"
	"net/http"
)

func (h Handler) UpdateFlag(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	req, err := ioutil.ReadAll(r.Body)
	if err != nil {
		msg := fmt.Sprintf("unable to read request body: %v", err)
		utils.ErrLog.Printf("%s", msg)
		utils.ErrorResponse(w, r, http.StatusBadRequest, msg)
		return
	}

	id, err := h.idFromParams(w, r, "flag")
	if err != nil {
		return
	}

	pub, res, err := h.DM.UpdateFlag(&req, id)
	h.ProcessServices(pub, "flag-update-channel")
	h.ComposeResponse(w, r, res, err)
}

func (h Handler) ToggleFlag(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	req, err := ioutil.ReadAll(r.Body)
	if err != nil {
		msg := fmt.Sprintf("unable to read request body: %v", err)
		utils.ErrLog.Printf("%s", msg)
		utils.ErrorResponse(w, r, http.StatusBadRequest, msg)
		return
	}

	id, err := h.idFromParams(w, r, "flag")
	if err != nil {
		return
	}

	pub, res, err := h.DM.ToggleFlag(&req, id)

	h.ProcessServices(pub, "flag-toggle-channel")

	h.ComposeResponse(w, r, res, err)
}

func (h Handler) UpdateAudience(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	req, err := ioutil.ReadAll(r.Body)
	if err != nil {
		msg := fmt.Sprintf("unable to read request body: %v", err)
		utils.ErrLog.Printf("%s", msg)
		utils.ErrorResponse(w, r, http.StatusBadRequest, msg)
		return
	}

	id, err := h.idFromParams(w, r, "flag")
	if err != nil {
		return
	}

	pub, res, err := h.DM.UpdateAudience(&req, id)

	h.ProcessServices(pub, "flag-update-channel")

	h.ComposeResponse(w, r, res, err)
}
