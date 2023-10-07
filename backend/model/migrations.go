package model

import "gorm.io/gorm"

func Migrate(DB *gorm.DB) {
	DB.AutoMigrate(
		&User{},
		&Chatpdf{},
		&ChatpdfMessage{},
	)
}
