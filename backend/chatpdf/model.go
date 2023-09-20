package chatpdf

import (
	"cb/common"
	"cb/users"
)

type Chatpdf struct {
	common.Model
	Name   string     `json:"name" gorm:"not null" validate:"required"`
	Key    string     `json:"key" gorm:"not null" validate:"required"`
	UserID string     `json:"userId" gorm:"not null" validate:"required"`
	User   users.User `json:"user" validate:"-"`
}

type ChatpdfMessage struct {
	common.Model
	Content   string  `json:"content" gorm:"not null" validate:"required"`
	ChatpdfID uint    `json:"chatpdfId" gorm:"not null" validate:"required"`
	Chatpdf   Chatpdf `json:"chatpdf"`
	Role      string  `json:"role" gorm:"not null" validate:"required;oneof:system user"`
}
