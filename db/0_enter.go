package db

import (
	"gorm.io/gorm"
)

var modelList = []interface{}{
	&Config{},
	&Server{},
	&Map{},
}

func GetModelList() []interface{} {
	return modelList
}

func RegisterModel(db *gorm.DB) error {
	return db.AutoMigrate(modelList...)
}

func InitData() {
	GetMapInit()()
}
