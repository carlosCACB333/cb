package router

import (
	"github.com/carlosCACB333/cb-back/handler"
	"github.com/carlosCACB333/cb-back/server"

	"github.com/gofiber/fiber/v2"
)

func UserRouter(s *server.Server, r fiber.Router) {
	r.Post("/", handler.CreateUser(s))
	r.Put("/", handler.UpdateUser(s))
	r.Delete("/:id", handler.DeleUser(s))
}
