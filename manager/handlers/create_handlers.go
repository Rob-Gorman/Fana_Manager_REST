package handlers

import (
	"io/ioutil"
	"net/http"
)

func (h *Handler) CreateFlag(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	req, err := ioutil.ReadAll(r.Body)

	res, err := h.DM.CreateFlag(&req)
	h.ComposeResponse(w, r, res, err)

	h.ProcessServices(res, "flag-update-channel")
}

func (h *Handler) CreateAttribute(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	req, err := ioutil.ReadAll(r.Body)

	res, err := h.DM.CreateAttribute(&req)
	h.ComposeResponse(w, r, res, err)

	h.ProcessServices(res, "")
}

func (h *Handler) CreateAudience(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	req, err := ioutil.ReadAll(r.Body)

	res, err := h.DM.CreateAudience(&req)
	h.ComposeResponse(w, r, res, err)

	h.ProcessServices(res, "audience-update-channel")
}
