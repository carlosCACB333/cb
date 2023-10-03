package main

import (
	"cb/chatpdf"
	"cb/middleware"
	"cb/users"
	"os"

	"github.com/gofiber/fiber/v2"
)

func SetupRouter(app *fiber.App) {

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("home", fiber.Map{
			"Title": "carloscb",
		}, "layouts/main")
	})

	{ // API V1

		{ // CONECTION FRONT TO BACK
			v1Public := app.Group("/api/v1/public")
			v1Public.Use(middleware.ApiKeyMiddleware(os.Getenv("X_API_KEY_PUBLIC")))
			{ //NOT REQUIRED AUTHENTICATION
				users.AuthRoutes(v1Public.Group("/auth"))
			}
			{ //REQUIRED AUTHENTICATION
				v1Public.Use(middleware.AuthMiddleware())
				v1Public.Get("/resource/:id", chatpdf.GetChatFile)
				v1Public.Post("/chatpdf", chatpdf.CreateChat)

			}
		}

		{ // CONECTION ONLY BACK TO BACK
			v1Private := app.Group("/api/v1")
			v1Private.Use(middleware.ApiKeyMiddleware(os.Getenv("X_API_KEY")))
			{ //NOT REQUIRED AUTHENTICATION
				users.UserRoutes(v1Private.Group("/user"))
			}
			{ //REQUIRED AUTHENTICATION
				v1Private.Use(middleware.AuthMiddleware())
				chatpdf.ChatPdfRoutes(v1Private.Group("/chatpdf"))
				chatpdf.BootRoutesRoutes(v1Private.Group("/boot"))
				chatpdf.MessagesRoutes(v1Private.Group("/message"))
			}
		}

	}

}
