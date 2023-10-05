package main

import (
	"cb/utils"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/django/v3"
)

func main() {

	if os.Getenv("STAGE") == "development" {
		// InitMigrations()
	}

	engine := django.New("./views", ".html")
	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		AppName:       "Chatbot API",
		Views:         engine,
		ErrorHandler:  utils.ErrorHandler,
	})
	app.Use(logger.New())
	app.Static("/public", "./public")

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000,https://carloscb.com,https://www.carloscb.com",
		AllowHeaders: "Origin, Content-Type, Authorization, x-api-key",
	}))

	SetupRouter(app)
	fmt.Println("🚀 Starting server on port: " + os.Getenv("PORT"))
	log.Fatal(app.Listen(":" + os.Getenv("PORT")))

}
