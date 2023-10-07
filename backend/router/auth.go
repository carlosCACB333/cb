package router

import (
	"cb/handler"
	"cb/server"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(s *server.Server, r fiber.Router) {
	r.Post("/register", handler.AuthRegister(s))
	r.Post("/login", handler.AuthLogin(s))
	r.Post("/change-password", handler.ChangePassword(s))
}
