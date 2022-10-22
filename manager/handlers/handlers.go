package handlers

import (
	dm "manager/data/datamodel"
	pub "manager/publisher"
	"manager/utils"
	"net/http"
)

type Handler struct {
	DM  *dm.DataModel
	Pub *pub.Pub
}

func New(dm *dm.DataModel, p *pub.Pub) *Handler {
	return &Handler{dm, p}
}

func (h *Handler) ComposeResponse(w http.ResponseWriter, r *http.Request, res *[]byte, err error) {
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	utils.CreatedResponse(w, r, &res)
}

// services extracted to one location
func (h *Handler) ProcessServices(res *[]byte, channel string) {
	if channel != "" {
		h.Pub.PublishContent(res, channel)
	}
	h.DM.RefreshCache()
}
