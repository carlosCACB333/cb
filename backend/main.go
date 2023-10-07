package main

import (
	"cb/router"
	"cb/server"
	"cb/util"
	"context"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/django/v3"
	"gorm.io/driver/postgres"
)

func main() {

	server := server.New(
		context.Background(),
		server.Config{
			Port: ":" + os.Getenv("PORT"),
			Config: fiber.Config{
				CaseSensitive: true,
				AppName:       "Chatbot API",
				Views:         django.New("./view", ".html"),
				ErrorHandler:  util.ErrorHandler,
			},
			Dialector:    postgres.Open(os.Getenv("DB_URL")),
			ApiKey:       os.Getenv("X_API_KEY"),
			ApiKeyPublic: os.Getenv("X_API_KEY_PUBLIC"),
			JWTSecret:    os.Getenv("JWT_SECRET"),
		})
	// server.Migrations()
	server.Start(router.SetupRouter)

}
