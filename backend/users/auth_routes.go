package users

import (
	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(r fiber.Router) {
	r.Post("/register", AuthRegister)
	r.Post("/login", AuthLogin)
	r.Post("/change-password", ChangePassword)
}
