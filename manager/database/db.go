package database

import (
	"manager/configs"
	"manager/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() *gorm.DB {
	dbUri := configs.DBConnStr()
	var err error
	DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  dbUri,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	utils.HandleErr(err, "Could not connect to DB at URI")

	return DB
}
