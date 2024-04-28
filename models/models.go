package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	Data string
}

func MigrateItems(db *gorm.DB) error {
	err := db.AutoMigrate(&Item{})
	return err
}
