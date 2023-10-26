package model

import "time"

type Session struct {
	Model
	UserID    string    `json:"user_id"`
	User      *User     `gorm:"foreignKey:UserID"`
	Token     string    `json:"token"`
	UserAgent string    `json:"user_agent"`
	IP        string    `json:"ip"`
	ExpiresAt time.Time `json:"expires_at"`
	IsBlocked bool      `json:"is_blocked"`
}
