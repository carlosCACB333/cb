package model

import "gorm.io/gorm"

func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		&User{},
		&Session{},
		&Post{},
		&Category{},
		&Tag{},
		&Comment{},
		&Verification{},
	)
}
