package chatpdf

import (
	"cb/libs"
	"cb/users"
	"cb/utils"
	"sort"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateMessage(c *fiber.Ctx) error {
	chatId := c.Params("id")
	chat := Chatpdf{}
	db := libs.DBInit()
	user := c.Locals("user").(*users.User)
	if err := db.Where("id = ? AND user_id = ?", chatId, user.ID).First(&chat).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(utils.ResponseMsg(
			"error", "Chat no encontrado",
		))

	}

	var message ChatpdfMessage
	if err := c.BodyParser(&message); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ResponseMsg(
			"error", "Datos incorrectos",
		))
	}
	message.ChatpdfID = chat.ID
	message.ID = uuid.New().String()
	if err := utils.ValidateFields(message); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.Response(
			"error", "Datos incorrectos", err, nil,
		))
	}
	if err := db.Create(&message).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ResponseMsg(
			"error", "Error al crear mensaje",
		))
	}
	return c.JSON(utils.Response(
		"success", "Mensaje creado correctamente",
		message,
		nil,
	))

}

func GetLastMessages(c *fiber.Ctx) error {
	db := libs.DBInit()
	chatId := c.Params("id")
	user := c.Locals("user").(*users.User)
	var chat Chatpdf
	var messages []ChatpdfMessage

	// validate if chat exists and belongs to user
	if err := db.Where("id = ? AND user_id = ?", chatId, user.ID).First(&chat).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(utils.ResponseMsg(
			"error", "Chat no encontrado",
		))
	}

	if err := db.Where("chatpdf_id = ?", chat.ID).Order("created_at desc").Limit(6).Find(&messages).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(utils.ResponseMsg(
			"error", "Mensajes no encontrados",
		))
	}

	sort.Slice(messages, func(i, j int) bool {
		return messages[i].CreatedAt.Before(messages[j].CreatedAt)
	})

	return c.JSON(utils.Response(
		"success", "Mensajes encontrados",
		messages,
		nil,
	))

}
