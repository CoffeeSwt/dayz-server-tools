package model

import "gorm.io/gorm"

var modelList = []interface{}{
	&Config{},
}

func RegisterModel(db *gorm.DB) error {
	return db.AutoMigrate(modelList...)
}
