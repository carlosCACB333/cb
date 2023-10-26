package model

import "time"

type Verification struct {
	Model
	User      User      `json:"user" gorm:"not null" validate:"-"`
	UserId    string    `json:"userId" gorm:"not null" validate:"required"`
	Code      string    `json:"code" gorm:"not null" validate:"required"`
	ExpiresAt time.Time `json:"expiresAt" gorm:"not null;default:now() + interval '10 minutes'" validate:"required"`
	IsUsed    bool      `json:"isUsed" gorm:"not null;default:false" validate:"required"`
}
