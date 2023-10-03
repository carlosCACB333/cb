package users

import (
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(r fiber.Router) {
	r.Post("/", CreateUser)
	r.Put("/", UpdateUser)
	r.Delete("/:id", DeleUser)
}
