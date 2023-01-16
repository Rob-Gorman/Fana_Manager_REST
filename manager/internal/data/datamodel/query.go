package datamodel

import (
	"manager/internal/data/models"
	"manager/utils"
)

func (d *DataModel) GetAllFlags() (*[]byte, error) {
	var flags *[]models.Flag
	err := d.DB.Find(&flags).Error

	if err != nil {
		utils.ErrLog.Printf("DB can't find flags %v", err)
		return nil, err
	}

	res := models.AllFlagsRes(flags)
	return models.ToJSON(res)
}

func (d *DataModel) GetAllAudiences() (*[]byte, error) {
	auds := []models.Audience{}

	err := d.DB.Preload("Conditions").Find(&auds).Error
	if err != nil {
		return nil, err
	}

	res := models.AllAudsRes(&auds)
	return models.ToJSON(res)
}

func (d *DataModel) GetAllAttributes() (*[]byte, error) {
	var attrs []models.Attribute

	err := d.DB.Find(&attrs).Error
	if err != nil {
		return nil, err
	}

	return models.ToJSON(attrs)
}

func (d *DataModel) GetFlag(id int) (*[]byte, error) {
	var flag models.Flag
	auds := []models.AudienceNoCondsResponse{}

	err := d.DB.Preload("Audiences").First(&flag, id).Error
	if err != nil {
		utils.ErrLog.Printf("could not find flag %d: %v", id, err)
		return nil, err
	}

	for ind := range flag.Audiences {
		auds = append(auds, models.AudienceNoCondsResponse{Audience: &flag.Audiences[ind]})
	}

	res := &models.FlagResponse{Flag: &flag, Audiences: auds}
	return models.ToJSON(res)
}

func (d *DataModel) GetAudience(id int) (*[]byte, error) {
	var aud models.Audience

	err := d.DB.Preload("Flags").Preload("Conditions").First(&aud, id).Error

	if err != nil {
		utils.ErrLog.Printf("could not find audience %d: %v", id, err)
		return nil, err
	}

	conds := d.GetEmbeddedConds(aud)       // can we do the DB work elsewhere?
	flags := d.GetEmbeddedFlags(aud.Flags) // no db calls - move to Models?

	res := models.AudienceResponse{
		Audience:   &aud,
		Conditions: conds,
		Flags:      flags,
	}

	return models.ToJSON(res)
}

func (d *DataModel) GetAttribute(id int) (*[]byte, error) {
	var attr models.Attribute

	err := d.DB.Preload("Conditions").First(&attr, id).Error

	if err != nil {
		utils.ErrLog.Printf("could not find attribute %d: %v", id, err)
		return nil, err
	}

	res := d.BuildAttrResponse(attr)
	return models.ToJSON(res)
}

func (d *DataModel) GetAuditLogs() (*[]byte, error) {
	flags := []models.FlagLog{}
	d.DB.Find(&flags)

	auds := []models.AudienceLog{}
	d.DB.Find(&auds)

	attrs := []models.AttributeLog{}
	d.DB.Find(&attrs)

	res := models.AuditResponse{
		FlagLogs:      flags,
		AudienceLogs:  auds,
		AttributeLogs: attrs,
	}

	return models.ToJSON(res)
}

func (d *DataModel) GetSdkKeys() (*[]byte, error) {
	sdks := []models.Sdkkey{}
	d.DB.Find(&sdks)
	return models.ToJSON(&sdks)
}
