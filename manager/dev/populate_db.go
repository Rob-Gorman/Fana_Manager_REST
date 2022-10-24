package dev

import (
	"manager/data/datamodel"
	"manager/data/models"
)

func RefreshSchema(d *datamodel.DataModel) {
	db := d.DB
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
	SeedDB(d)
}
