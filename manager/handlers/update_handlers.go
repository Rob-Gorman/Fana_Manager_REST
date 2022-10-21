package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"manager/models"
	"manager/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func (h Handler) UpdateFlag(w http.ResponseWriter, r *http.Request) {
	var flagReq models.FlagSubmit

	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)

	err := json.Unmarshal(body, &flagReq)
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Invalid resource ID."))
		return
	}

	fr := FlagReqToFlag(flagReq, h)
	var flag models.Flag
	h.DB.First(&flag, id)
	flag.Audiences = fr.Audiences
	flag.DisplayName = fr.DisplayName
	flag.Sdkkey = fr.Sdkkey

	if flagReq.Audiences != nil {
		h.DB.Model(&flag).Omit("Audiences.*").Association("Audiences").Replace(flag.Audiences)
	}

	err = h.DB.Omit("Audiences").Session(&gorm.Session{
		SkipHooks:            true,
	}).Updates(&flag).Error
	

	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	response := FlagToFlagResponse(flag, h)

	pub := FlagUpdateForPublisher(h.DB, []models.Flag{flag})
	PublishContent(&pub, "flag-update-channel")
	RefreshCache(h.DB)

	utils.UpdatedResponse(w, r, &response)
}

func (h Handler) ToggleFlag(w http.ResponseWriter, r *http.Request) {
	fmt.Println("\nGot a flag toggle!")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Invalid resource ID."))
		return
	}

	togglef := struct {
		Status bool `json:"status"`
	}{}

	body, _ := ioutil.ReadAll(r.Body)

	err = json.Unmarshal(body, &togglef)

	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	var flag models.Flag
	h.DB.Find(&flag, id)
	flag.Status = togglef.Status
	flag.DisplayName = fmt.Sprintf("__%v", flag.Status) // hacky way to clue it's a toggle action, see flag update hook
	err = h.DB.Select("status").Updates(&flag).Error
	if err != nil {
		utils.NoRecordResponse(w, r, err)
		return
	}
	h.DB.First(&flag, id)
	response := models.FlagNoAudsResponse{Flag: &flag}
	pub := FlagUpdateForPublisher(h.DB, []models.Flag{flag})
	PublishContent(&pub, "flag-toggle-channel")

	RefreshCache(h.DB)
	
	utils.UpdatedResponse(w, r, &response)
}

func (h Handler) UpdateAudience(w http.ResponseWriter, r *http.Request) {
	var req models.Audience

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Invalid resource ID."))
		return
	}

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	err = json.Unmarshal(body, &req)
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	aud := BuildAudUpdate(req, id, h)

	if req.Conditions != nil {
		h.DB.Model(&aud).Association("Conditions").Replace(aud.Conditions)
	}

	err = h.DB.Session(&gorm.Session{
		FullSaveAssociations: true,
		SkipHooks: true,
		}).Updates(&aud).Error
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	h.DB.Model(&models.Audience{}).Preload("Flags").Preload("Conditions").Find(&aud)

	response := models.AudienceResponse{
		Audience:   &aud,
		Conditions: GetEmbeddedConds(aud, h.DB),
		Flags:      GetEmbeddedFlags(aud.Flags),
	}

	pub := FlagUpdateForPublisher(h.DB, aud.Flags)
	PublishContent(&pub, "flag-update-channel")
	RefreshCache(h.DB)

	utils.CreatedResponse(w, r, &response)
}
