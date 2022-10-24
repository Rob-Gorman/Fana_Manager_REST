package handlers

import (
	"manager/cache"
	"net/http"
)

func (h Handler) GetFlagset(w http.ResponseWriter, r *http.Request) {
	res, err := h.DM.GetFlagset()
	h.ComposeResponse(w, r, res, err)

	flagCache := cache.InitFlagCache()
	flagCache.FlushAllAsync()
	flagCache.Set("data", res)
}
