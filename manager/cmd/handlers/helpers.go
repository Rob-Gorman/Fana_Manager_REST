package handlers

import (
	"fmt"
	"manager/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (h Handler) idFromParams(w http.ResponseWriter, r *http.Request, relation string) (id int, err error) {
	idParam := mux.Vars(r)["id"]
	id, err = strconv.Atoi(idParam)

	if err != nil {
		msg := fmt.Sprintf("invalid %s id param: %d", relation, id)
		utils.ErrLog.Printf(msg)
		utils.ErrorResponse(w, r, http.StatusBadRequest, msg)
	}

	return
}
