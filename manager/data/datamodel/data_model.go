package datamodel

import (
	"manager/data"

	"gorm.io/gorm"
)

type DataModel struct {
	DB *gorm.DB
}

func New() (*DataModel, error) {
	db, err := data.Init()
	return &DataModel{db}, err
}