package datamodel

import (
	"manager/data/models"
	"manager/utils"

	"gorm.io/gorm"
)

func (d *DataModel) CreateFlag(req *[]byte) (*[]byte, error) {
	var flagReq models.FlagSubmit

	err := flagReq.FromJSON(req)
	if err != nil {
		return nil, err
	}

	flag := d.FlagReqToFlag(flagReq) // can refactor to models?

	err = d.DB.Omit("Audiences.*").Session(&gorm.Session{FullSaveAssociations: true}).Create(&flag).Error
	if err != nil {
		return nil, utils.DuplicateError(err)
	}

	d.DB.Preload("Audiences").Find(&flag)
	respAuds := []models.AudienceNoCondsResponse{}

	for ind := range flag.Audiences {
		respAuds = append(respAuds, models.AudienceNoCondsResponse{
			Audience: &flag.Audiences[ind],
		})
	}

	res := models.FlagResponse{
		Flag:      &flag,
		Audiences: respAuds,
	}

	return models.ToJSON(res)
}

func (d *DataModel) CreateAttribute(req *[]byte) (*[]byte, error) {
	var attrReq models.Attribute

	err := attrReq.FromJSON(req)
	if err != nil {
		return nil, err
	}

	err = d.DB.Create(&attrReq).Error
	if err != nil {
		return nil, utils.DuplicateError(err)
	}

	d.DB.Find(&attrReq)

	return models.ToJSON(attrReq)
}

func (d *DataModel) CreateAudience(req *[]byte) (*[]byte, error) {
	var aud models.Audience

	err := aud.FromJSON(req)
	if err != nil {
		return nil, err
	}

	err = d.DB.Session(&gorm.Session{FullSaveAssociations: true}).Create(&aud).Error
	if err != nil {
		return nil, utils.DuplicateError(err)
	}

	d.DB.Model(&models.Audience{}).Preload("Conditions").Find(&aud)

	// can we refactor?
	conds := d.GetEmbeddedConds(aud)

	res := aud.ToResponse(conds)

	return models.ToJSON(res)
}
