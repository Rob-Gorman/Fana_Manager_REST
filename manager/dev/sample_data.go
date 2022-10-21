package dev

import (
	"manager/handlers"
	"manager/models"

	"gorm.io/gorm"
)

func SeedDB(db *gorm.DB) {
	seedFlags(db)
	seedAttributes(db)
	seedAudiences(db)
	seedSdks(db)
	handlers.BuildFlagset(db)
}

func seedSdks(db *gorm.DB) {
	key1 := handlers.NewSDKKey("***-*****-**")
	key2 := handlers.NewSDKKey("***-*****-**")
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
