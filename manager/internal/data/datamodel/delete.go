package datamodel

import (
	"manager/internal/data/models"
	"manager/utils"
	"net/http"
)

func (d *DataModel) DeleteFlag(id int) (int, error) {
	flag := &models.Flag{}
	err := d.DB.Preload("Audiences").First(&flag, id).Error
	if err != nil {
		utils.ErrLog.Printf("%v", err)
		return utils.NotFoundErr("flag", id)
	}

	err = d.DB.Model(&flag).Association("Audiences").Delete(flag.Audiences)
	if err != nil {
		msg := "unable to delete flag audiences"
		utils.ErrLog.Printf("%s: %v", msg, err)
		return utils.UnprocessableErr(msg)
	}

	err = d.DB.Unscoped().Delete(&flag).Error
	if err != nil {
		msg := "unable to delete flag"
		utils.ErrLog.Printf("%s: %v", msg, err)
		return utils.UnprocessableErr(msg)
	}
	return 200, nil
}

func (d *DataModel) DeleteAudience(id int) (int, error) {
	aud := &models.Audience{}
	err := d.DB.Preload("Flags").First(&aud, id).Error
	if err != nil {
		utils.ErrLog.Printf("%v", err)
		return utils.NotFoundErr("audience", id)
	}

	if !OrphanedAud(aud) {
		msg := "unable to delete audiences assigned to flags"
		utils.ErrLog.Printf(msg)
		return utils.UnprocessableErr(msg)
	}

	d.DB.Model(&aud).Association("Flags").Delete(aud.Flags)
	err = d.DB.Unscoped().Delete(&aud).Error
	if err != nil {
		msg := "unable to delete audience"
		utils.ErrLog.Printf("%s: %v", msg, err)
		return utils.UnprocessableErr(msg)
	}
	return 200, nil
}

func (d *DataModel) DeleteAttribute(id int) (int, error) {
	attr := &models.Attribute{}
	err := d.DB.Preload("Conditions").First(attr, id).Error
	if err != nil {
		utils.ErrLog.Printf("%v", err)
		return utils.NotFoundErr("attribute", id)
	}

	if !OrphanedAttr(attr) {
		msg := "Cannot delete attribute assigned to audiences"
		utils.ErrLog.Printf(msg)
		return utils.UnprocessableErr(msg)
	}

	err = d.DB.Unscoped().Delete(&attr).Error
	if err != nil {
		msg := "unable to delete audience"
		utils.ErrLog.Printf("%s: %v", msg, err)
		return utils.UnprocessableErr(msg)
	}

	return 200, nil
}

func (d *DataModel) RegenSDKkey(id int) (*[]byte, int, error) {
	sdk := models.Sdkkey{}

	err := d.DB.Find(&sdk, id).Error
	if err != nil {
		utils.ErrLog.Printf("%v", err)
		code, err := utils.NotFoundErr("sdk key", id)
		return nil, code, err
	}

	newSDK := models.Sdkkey{
		Key:  NewSDKKey(sdk.Key),
		Type: sdk.Type,
	}

	err = d.DB.Create(&newSDK).Error
	if err != nil {
		msg := "unable to create new SDK key"
		utils.ErrLog.Printf("%s: %v", msg, err)
		code, err := utils.InternalErr(msg)
		return nil, code, err
	}

	d.DB.Unscoped().Delete(&sdk)

	d.DB.Find(&newSDK)

	res, err := models.ToJSON(newSDK)
	return res, http.StatusCreated, err
}
