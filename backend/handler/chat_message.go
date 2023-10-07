package handler

import (
	"cb/model"
	"cb/server"
	"cb/util"
	"sort"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateMessage(s *server.Server) fiber.Handler {
	return func(c *fiber.Ctx) error {
		chatId := c.Params("id")
		chat := model.Chatpdf{}
		db := s.DB()
		user := c.Locals("user").(*model.User)
		if err := db.Where("id = ? AND user_id = ?", chatId, user.ID).First(&chat).Error; err != nil {
			return util.NewError(fiber.StatusNotFound, "Chat no encontrado", nil)
		}

		var message model.ChatpdfMessage
		if err := c.BodyParser(&message); err != nil {
			return util.NewError(fiber.StatusBadRequest, "Datos incorrectos", nil)
		}
		message.ChatpdfID = chat.ID
		message.ID = uuid.New().String()
		if err := util.ValidateFields(message); err != nil {
			return util.NewError(fiber.StatusBadRequest, "Datos incorrectos", err)
		}
		if err := db.Create(&message).Error; err != nil {
			return util.NewError(fiber.StatusBadRequest, "Error al crear mensaje", nil)

		}

		return c.JSON(util.NewBody(util.Body{
			Data:    message,
			Message: "Mensaje creado correctamente",
		}))

	}
}

func GetLastMessages(s *server.Server) fiber.Handler {
	return func(c *fiber.Ctx) error {
		db := s.DB()
		chatId := c.Params("id")
		user := c.Locals("user").(*model.User)
		var chat model.Chatpdf
		var messages []model.ChatpdfMessage

		// validate if chat exists and belongs to user
		if err := db.Where("id = ? AND user_id = ?", chatId, user.ID).First(&chat).Error; err != nil {
			return util.NewError(fiber.StatusNotFound, "Chat no encontrado", nil)
		}

		if err := db.Where("chatpdf_id = ?", chat.ID).Order("created_at desc").Limit(6).Find(&messages).Error; err != nil {
			return util.NewError(fiber.StatusNotFound, "Mensajes no encontrados", nil)

		}

		sort.Slice(messages, func(i, j int) bool {
			return messages[i].CreatedAt.Before(messages[j].CreatedAt)
		})

		return c.JSON(util.NewBody(util.Body{
			Data: messages,
		}))

	}
}
