package router

import (
	"cb/handler"
	"cb/server"

	"github.com/gofiber/fiber/v2"
)

func ChatPdfRoutes(s *server.Server, r fiber.Router) {
	r.Get("", handler.GetAllChats(s))
	r.Get("/:id", handler.GetChatsById(s))
	r.Delete("/:id", handler.DeleteChat(s))
}
func BootRoutesRoutes(s *server.Server, r fiber.Router) {
	r.Post("/:id", handler.GetSimilarity(s))
}

func MessagesRoutes(s *server.Server, r fiber.Router) {
	r.Post("/:id/new", handler.CreateMessage(s))
	r.Get("/:id/last", handler.GetLastMessages(s))
}
