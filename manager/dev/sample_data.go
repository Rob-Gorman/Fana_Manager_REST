package dev

import (
	"manager/internal/data/datamodel"
	"manager/internal/data/models"

	"gorm.io/gorm"
)

func SeedDB(d *datamodel.DataModel) {
	db := d.DB
	seedFlags(db)
	seedAttributes(db)
	seedAudiences(db)
	seedSdks(db)
	d.BuildFlagset()
}

func seedSdks(db *gorm.DB) {
	key1 := datamodel.NewSDKKey("***-*****-**")
	key2 := datamodel.NewSDKKey("***-*****-**")
	var sdkkeys = []models.Sdkkey{
		{Key: key1},
		{Key: key2, Type: "server"},
	}
	db.Create(&sdkkeys)
}

func seedFlags(db *gorm.DB) {
	var flags = []models.Flag{
		{Key: "sample_flag", DisplayName: "Sample Flag"},
	}
	db.Create(&flags)
}

func seedAttributes(db *gorm.DB) {
	var attrs = []models.Attribute{
		{Key: "sample_attribute", Type: "STR", DisplayName: "Sample Attribute"},
	}
	db.Create(&attrs)
}

func seedAudiences(db *gorm.DB) {
	sample := models.Audience{
		Key:         "sample_audience",
		DisplayName: "Sample Audience",
		Conditions:  []models.Condition{},
	}

	// db.Create(&ca_stu) // could just do this for a single audience
	auds := []models.Audience{sample}
	db.Create(&auds)
}
