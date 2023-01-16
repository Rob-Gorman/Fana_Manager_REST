package handlers

import (
	dm "manager/internal/data/datamodel"
	pub "manager/internal/publisher"
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

	utils.PayloadResponse(w, r, res)
}

// services extracted to one location
func (h *Handler) ProcessServices(res *[]byte, topic string) {
	if topic != "" {
		h.Pub.PublishContent(res, topic)
	}
	h.DM.RefreshCache()
}
