package chatpdf

import (
	"github.com/gofiber/fiber/v2"
)

func ChatPdfRoutes(r fiber.Router) {
	r.Get("", GetAllChats)
	r.Get("/:id", GetChatsById)
	r.Delete("/:id", DeleteChat)
}
func BootRoutesRoutes(r fiber.Router) {
	r.Post("/:id", GetSimilarity)
}

func MessagesRoutes(r fiber.Router) {
	r.Post("/:id/new", CreateMessage)
	r.Get("/:id/last", GetLastMessages)
}
