package model

import "time"

type Model struct {
	ID        string    `json:"id" gorm:"primary_key;default:null"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime:true"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime:true"`
}
