package utils

import (
	"github.com/gofiber/fiber/v2"
)

func Response(status string, message string, data interface{}, others fiber.Map) fiber.Map {
	res := fiber.Map{
		"status":  status,
		"message": message,
		"data":    data,
	}
	for k, v := range others {
		res[k] = v
	}

	return res
}

func ResponseMsg(status string, message string) fiber.Map {
	return Response(status, message, nil, nil)
}
