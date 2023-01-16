package data

import (
	"fmt"
	"manager/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() (*gorm.DB, error) {
	dbUri := dbConnStr()

	DB, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dbUri,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		utils.ErrLog.Printf("DB cannot connect %v", err)
	}

	return DB, err
}

func dbConnStr() string {
	variables := utils.GetEnvVars("DB_HOST", "DB_USER", "DB_NAME", "DB_PW", "DB_PORT")
	dbUri := fmt.Sprintf(
		"host=%s user=%s dbname=%s sslmode=disable password=%s port=%s",
		variables...,
	)

	return dbUri
}
