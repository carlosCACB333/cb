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
		return utils.NewError(fiber.StatusNotFound, "Chat no encontrado", nil)
	}

	var message ChatpdfMessage
	if err := c.BodyParser(&message); err != nil {
		return utils.NewError(fiber.StatusBadRequest, "Datos incorrectos", nil)
	}
	message.ChatpdfID = chat.ID
	message.ID = uuid.New().String()
	if err := utils.ValidateFields(message); err != nil {
		return utils.NewError(fiber.StatusBadRequest, "Datos incorrectos", err)
	}
	if err := db.Create(&message).Error; err != nil {
		return utils.NewError(fiber.StatusBadRequest, "Error al crear mensaje", nil)

	}

	return c.JSON(utils.NewBody(utils.Body{
		Data:    message,
		Message: "Mensaje creado correctamente",
	}))

}

func GetLastMessages(c *fiber.Ctx) error {
	db := libs.DBInit()
	chatId := c.Params("id")
	user := c.Locals("user").(*users.User)
	var chat Chatpdf
	var messages []ChatpdfMessage

	// validate if chat exists and belongs to user
	if err := db.Where("id = ? AND user_id = ?", chatId, user.ID).First(&chat).Error; err != nil {
		return utils.NewError(fiber.StatusNotFound, "Chat no encontrado", nil)
	}

	if err := db.Where("chatpdf_id = ?", chat.ID).Order("created_at desc").Limit(6).Find(&messages).Error; err != nil {
		return utils.NewError(fiber.StatusNotFound, "Mensajes no encontrados", nil)

	}

	sort.Slice(messages, func(i, j int) bool {
		return messages[i].CreatedAt.Before(messages[j].CreatedAt)
	})

	return c.JSON(utils.NewBody(utils.Body{
		Data: messages,
	}))

}
