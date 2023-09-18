package main

import (
	"cb/chatpdf"
	"cb/libs"
	"cb/users"
	"fmt"
)

func InitMigrations() {
	err := libs.DBInit().AutoMigrate(
		&users.User{},
		&chatpdf.Chatpdf{},
		&chatpdf.ChatpdfMessage{},
	)

	fmt.Println(err)
}
