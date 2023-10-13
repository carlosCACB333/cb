package router

import (
	"github.com/carlosCACB333/cb-back/handler"
	"github.com/carlosCACB333/cb-back/server"

	"github.com/gofiber/fiber/v2"
)

func AuthRouter(s *server.Server, r fiber.Router) {
	h := handler.NewAuthHandler(s)
	r.Post("/register", h.AuthRegister)
	r.Post("/login", h.AuthLogin)
	r.Post("/change-password", h.ChangePassword)
}
