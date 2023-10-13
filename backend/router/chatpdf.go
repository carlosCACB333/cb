package router

import (
	"github.com/carlosCACB333/cb-back/handler"
	"github.com/carlosCACB333/cb-back/server"

	"github.com/gofiber/fiber/v2"
)

func ChatPdfRouter(s *server.Server, r fiber.Router) {
	r.Get("", handler.GetAllChats(s))
	r.Get("/:id", handler.GetChatsById(s))
	r.Delete("/:id", handler.DeleteChat(s))
}
func BootRouter(s *server.Server, r fiber.Router) {
	r.Post("/:id", handler.GetSimilarity(s))
}

func MessagesRouter(s *server.Server, r fiber.Router) {
	r.Post("/:id/new", handler.CreateMessage(s))
	r.Get("/:id/last", handler.GetLastMessages(s))
}
