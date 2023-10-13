package middleware

import (
	"github.com/carlosCACB333/cb-back/util"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	if e, ok := err.(*fiber.Error); ok {
		return c.Status(e.Code).JSON(util.NewBody(util.Body{
			Status:  "error",
			Message: e.Message,
		}))
	} else if e, ok := err.(util.Error); ok {
		return c.Status(e.Code).JSON(util.NewBody(util.Body{
			Message: e.Message,
			Status:  "error",
			Data:    e.Data,
		}))

	}
	return c.Status(fiber.StatusInternalServerError).JSON(util.NewBody(util.Body{
		Status:  "error",
		Message: "Error interno del servidor",
	}))

}
