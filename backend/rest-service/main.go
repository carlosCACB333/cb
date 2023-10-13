package main

import (
	"context"
	"os"
	"time"

	"github.com/carlosCACB333/cb-back/middleware"
	"github.com/carlosCACB333/cb-back/router"
	"github.com/carlosCACB333/cb-back/server"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/django/v3"
	"gorm.io/driver/postgres"
)

func main() {

	s := server.New(
		context.Background(),
		server.Config{
			Config: fiber.Config{
				CaseSensitive: true,
				AppName:       "Chatbot API",
				Views:         django.New("./view", ".html"),
				ErrorHandler:  middleware.ErrorHandler,
				IdleTimeout:   time.Second * 20,
				ReadTimeout:   time.Second * 10,
				WriteTimeout:  time.Second * 10,
			},
			Port:         ":" + os.Getenv("PORT"),
			Dialector:    postgres.Open(os.Getenv("DB_URL")),
			ApiKey:       os.Getenv("X_API_KEY"),
			ApiKeyPublic: os.Getenv("X_API_KEY_PUBLIC"),
			JWTSecret:    os.Getenv("JWT_SECRET"),
			NatsUrl:      os.Getenv("NATS_URL"),
		})
	// server.Migrations()

	s.Start(router.SetupRouter)

	s.Shutdown()
}
