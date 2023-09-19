package users

import "cb/common"

type User struct {
	common.Model
	Username      string `json:"username" gorm:"" validate:"required"`
	FirstName     string `json:"firstName" gorm:"" validate:"required"`
	LastName      string `json:"lastName" gorm:"" validate:"required"`
	Email         string `json:"email" gorm:"not null;unique" validate:"required,email"`
	Gender        string `json:"gender" gorm:"" validate:"required"`
	Password      string `json:"password" gorm:"" validate:"required"`
	Photo         string `json:"photo" gorm:"" validate:"required"`
	Phone         string `json:"phone" gorm:"type:varchar(9)" validate:"required,numeric,len=9"`
	SocialAccount string `json:"socialAccount" gorm:""`
	Status        string `json:"status" gorm:"not null,default:'active'" `
}
