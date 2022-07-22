package handlers

import (
	"manager/utils"
	"net/http"
)

func (h Handler) GetFlagset(w http.ResponseWriter, r *http.Request) {
	fs := api.BuildFlagset(h.DB)
	utils.PayloadResponse(w, r, &fs)
}
