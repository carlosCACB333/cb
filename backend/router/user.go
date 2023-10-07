package router

import (
	"cb/handler"
	"cb/server"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(s *server.Server, r fiber.Router) {
	r.Post("/", handler.CreateUser(s))
	r.Put("/", handler.UpdateUser(s))
	r.Delete("/:id", handler.DeleUser(s))
}
