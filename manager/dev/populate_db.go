package dev

import (
	"manager/models"

	"gorm.io/gorm"
)

func RefreshSchema(db *gorm.DB) {
	// re-creates them with the defined schema
	// and seeds some data
	var tables []interface{}
	tables = append(tables,
		&models.Flag{},
		&models.Audience{},
		&models.Attribute{},
		&models.Condition{},
		&models.FlagLog{},
		&models.AudienceLog{},
		&models.AttributeLog{},
		&models.Sdkkey{},
	)

	// create all relevant tables
	db.AutoMigrate(tables...)
	// seed sample data
	SeedDB(db)
}
