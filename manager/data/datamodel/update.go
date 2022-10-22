package datamodel

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"manager/data/models"
	"manager/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func (d *DataModel) UpdateFlag(w http.ResponseWriter, r *http.Request) {
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

	fr := FlagReqToFlag(flagReq, d)
	var flag models.Flag
	d.DB.First(&flag, id)
	flag.Audiences = fr.Audiences
	flag.DisplayName = fr.DisplayName
	flag.Sdkkey = fr.Sdkkey

	if flagReq.Audiences != nil {
		d.DB.Model(&flag).Omit("Audiences.*").Association("Audiences").Replace(flag.Audiences)
	}

	err = d.DB.Omit("Audiences").Session(&gorm.Session{
		SkipHooks: true,
	}).Updates(&flag).Error

	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	response := FlagToFlagResponse(flag, d)

	pub := FlagUpdateForPublisher(d.DB, []models.Flag{flag})
	PublishContent(&pub, "flag-update-channel")
	RefreshCache(d.DB)

	utils.UpdatedResponse(w, r, &response)
}

func (d *DataModel) ToggleFlag(w http.ResponseWriter, r *http.Request) {
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
	d.DB.Find(&flag, id)
	flag.Status = togglef.Status
	flag.DisplayName = fmt.Sprintf("__%v", flag.Status) // hacky way to clue it's a toggle action, see flag update hook
	err = d.DB.Select("status").Updates(&flag).Error
	if err != nil {
		utils.NoRecordResponse(w, r, err)
		return
	}
	d.DB.First(&flag, id)
	response := models.FlagNoAudsResponse{Flag: &flag}
	pub := FlagUpdateForPublisher(d.DB, []models.Flag{flag})
	PublishContent(&pub, "flag-toggle-channel")

	RefreshCache(d.DB)

	utils.UpdatedResponse(w, r, &response)
}

func (d *DataModel) UpdateAudience(w http.ResponseWriter, r *http.Request) {
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

	aud := BuildAudUpdate(req, id, d)

	if req.Conditions != nil {
		d.DB.Model(&aud).Association("Conditions").Replace(aud.Conditions)
	}

	err = d.DB.Session(&gorm.Session{
		FullSaveAssociations: true,
		SkipHooks:            true,
	}).Updates(&aud).Error
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	d.DB.Model(&models.Audience{}).Preload("Flags").Preload("Conditions").Find(&aud)

	response := models.AudienceResponse{
		Audience:   &aud,
		Conditions: GetEmbeddedConds(aud, d.DB),
		Flags:      GetEmbeddedFlags(aud.Flags),
	}

	pub := FlagUpdateForPublisher(d.DB, aud.Flags)
	PublishContent(&pub, "flag-update-channel")
	RefreshCache(d.DB)

	utils.CreatedResponse(w, r, &response)
}
