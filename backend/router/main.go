package router

import (
	"cb/handler"
	"cb/middleware"
	"cb/model"

	"cb/server"

	"github.com/gofiber/fiber/v2"
)

func SetupRouter(s *server.Server) {

	s.App().Get("/", func(c *fiber.Ctx) error {
		return c.Render("home", fiber.Map{
			"Title": "carloscb",
		}, "layouts/main")
	})

	{ // API V1

		{ // CONECTION FRONT TO BACK
			v1Public := s.App().Group("/api/v1/public")
			v1Public.Use(middleware.ApiKeyMiddleware(s.Config().ApiKeyPublic))
			{ //NOT REQUIRED AUTHENTICATION
				AuthRoutes(s, v1Public.Group("/auth"))
				v1Public.Get("/test", func(c *fiber.Ctx) error {
					body := map[string]interface{}{}
					c.BodyParser(&body)
					s.Hub(handler.WS_TEST).Broadcast(
						model.Message{
							Type:    "test",
							Payload: body,
						},
						nil)
					return c.JSON(fiber.Map{
						"message": "test",
					})
				})
			}
			{ //REQUIRED AUTHENTICATION
				v1Public.Use(middleware.AuthMiddleware(s))
				v1Public.Get("/resource/:id", handler.GetChatFile(s))
				v1Public.Post("/chatpdf", handler.CreateChat(s))
				{ // WEBSOCKET
					v1PublicWS := v1Public.Group("/ws")
					v1PublicWS.Use(middleware.VerifySocketUpgrade)
					v1PublicWS.Get("", handler.Handler(s))
				}

			}
		}

		{ // CONECTION ONLY BACK TO BACK
			v1Private := s.App().Group("/api/v1")
			v1Private.Use(middleware.ApiKeyMiddleware(s.Config().ApiKey))
			{ //NOT REQUIRED AUTHENTICATION
				UserRoutes(s, v1Private.Group("/user"))
			}
			{ //REQUIRED AUTHENTICATION
				v1Private.Use(middleware.AuthMiddleware(s))
				ChatPdfRoutes(s, v1Private.Group("/chatpdf"))
				BootRoutesRoutes(s, v1Private.Group("/boot"))
				MessagesRoutes(s, v1Private.Group("/message"))
			}
		}

	}

}
