package model

type User struct {
	Model
	Username  string `json:"username" gorm:"" validate:"required"`
	FirstName string `json:"firstName" gorm:"" validate:"required"`
	LastName  string `json:"lastName" gorm:"" validate:"required"`
	Email     string `json:"email" gorm:"not null;unique" validate:"required,email"`
	Gender    string `json:"gender" gorm:"" validate:"required"`
	Password  string `json:"-" gorm:"" validate:"required"`
	Photo     string `json:"photo" gorm:"" validate:"required"`
	Phone     string `json:"phone" gorm:"type:varchar(9)" validate:"required,numeric,len=9"`
	Status    string `json:"status" gorm:"not null,default:'NOT_VALIDATED_EMAIL'" `
}
