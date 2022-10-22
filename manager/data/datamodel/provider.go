package datamodel

import (
	"manager/cache"
	"manager/utils"
	"net/http"
)


func (d *DataModel) GetFlagset(w http.ResponseWriter, r *http.Request) {
	fs := BuildFlagset(d.DB)
	utils.PayloadResponse(w, r, &fs)
	
	flagCache := cache.InitFlagCache()
	flagCache.FlushAllAsync()
	flagCache.Set("data", &fs)

}
