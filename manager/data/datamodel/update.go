package datamodel

import (
	"encoding/json"
	"fmt"
	"manager/data/models"
	"manager/utils"

	"gorm.io/gorm"
)

func (d *DataModel) UpdateFlag(req *[]byte, id int) (*[]byte, *[]byte, error) {
	var flagReq models.FlagSubmit

	err := json.Unmarshal(*req, &flagReq)
	if err != nil {
		utils.ErrLog.Printf("failed to unmarshal request: %v", err)
		return nil, nil, err
	}

	fr := d.FlagReqToFlag(flagReq)
	var toUpdate models.Flag
	d.DB.First(&toUpdate, id)
	toUpdate.Audiences = fr.Audiences
	toUpdate.DisplayName = fr.DisplayName
	toUpdate.Sdkkey = fr.Sdkkey

	if flagReq.Audiences != nil {
		d.DB.Model(&toUpdate).Omit("Audiences.*").Association("Audiences").Replace(toUpdate.Audiences)
	}

	err = d.DB.Omit("Audiences").Session(&gorm.Session{
		SkipHooks: true,
	}).Updates(&toUpdate).Error

	if err != nil {
		utils.ErrLog.Printf("failed to update flag: %v", err)
		return nil, nil, err
	}

	res := d.FlagToFlagResponse(toUpdate)
	pub := d.FlagUpdateForPublisher([]models.Flag{toUpdate})
	pubret, err := models.ToJSON(pub)
	resret, err := models.ToJSON(res)
	return pubret, resret, err
}

func (d *DataModel) ToggleFlag(req *[]byte, id int) (*[]byte, *[]byte, error) {
	togglef := struct {
		Status bool `json:"status"`
	}{}

	err := json.Unmarshal(*req, &togglef)
	if err != nil {
		utils.ErrLog.Printf("failed to unmarshal request: %v", err)
		return nil, nil, err
	}

	var flag models.Flag
	d.DB.Find(&flag, id)
	flag.Status = togglef.Status
	flag.DisplayName = fmt.Sprintf("__%v", flag.Status) // hacky way to clue it's a toggle action, see flag update hook
	err = d.DB.Select("status").Updates(&flag).Error

	if err != nil {
		utils.ErrLog.Printf("failed to toggle flag: %v", err)
		return nil, nil, err
	}

	d.DB.First(&flag, id)
	res := models.FlagNoAudsResponse{Flag: &flag}
	pub := d.FlagUpdateForPublisher([]models.Flag{flag})

	pubret, err := models.ToJSON(pub)
	resret, err := models.ToJSON(res)

	return pubret, resret, err
}

func (d *DataModel) UpdateAudience(req *[]byte, id int) (pub *[]byte, res *[]byte, err error) {
	var update models.Audience
	var existing models.Audience

	err = json.Unmarshal(*req, &update)
	if err != nil {
		utils.ErrLog.Printf("failed to unmarshal request: %v", err)
		return nil, nil, err
	}

	d.DB.Find(&existing, id)
	existing.Update(&update)

	if update.Conditions != nil {
		d.DB.Model(&existing).Association("Conditions").Replace(existing.Conditions)
	}

	err = d.DB.Session(&gorm.Session{
		FullSaveAssociations: true,
		SkipHooks:            true,
	}).Updates(&existing).Error
	if err != nil {
		utils.ErrLog.Printf("failed to update audience: %v", err)
		return nil, nil, err
	}

	d.DB.Model(&models.Audience{}).Preload("Flags").Preload("Conditions").Find(&existing)

	aud := models.AudienceResponse{
		Audience:   &existing,
		Conditions: d.GetEmbeddedConds(existing),
		Flags:      d.GetEmbeddedFlags(existing.Flags),
	}

	pubmap := d.FlagUpdateForPublisher(existing.Flags)
	pub, err = models.ToJSON(pubmap)
	res, err = models.ToJSON(aud)

	return
}
