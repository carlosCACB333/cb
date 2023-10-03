package chatpdf

import (
	"cb/common"
	"cb/libs"
	"cb/users"
	"fmt"
	"strings"

	"cb/utils"
	"io"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func getUrl() string {
	return "https://" + os.Getenv("AWS_S3_BUCKET") + ".s3.amazonaws.com/"
}

func CreateChat(c *fiber.Ctx) error {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(utils.ResponseMsg(
			"error", "Error al obtener archivo",
		))
	}

	if fileHeader.Size > 1025*1024*10 {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ResponseMsg(
			"error", "El archivo no puede ser mayor a 10MB",
		))
	}
	fileName := fileHeader.Filename
	ID := uuid.New().String()
	split := strings.Split(fileName, ".")
	fileKey := ID + "." + split[len(split)-1]

	er := utils.PutObject(fileHeader, os.Getenv("AWS_S3_BUCKET"), fileKey)
	if er != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ResponseMsg(
			"error", "Error al subir archivo",
		))
	}
	user := c.Locals("user").(*users.User)

	chat := Chatpdf{
		Model:       common.Model{ID: ID},
		Name:        fileName,
		Key:         fileKey,
		UserID:      user.ID,
		ContentType: fileHeader.Header.Get("Content-Type"),
	}

	if err := libs.DBInit().Create(&chat).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ResponseMsg(
			"error", "Error al crear chat",
		))
	}
	// VECTORIZE AND SAVE ON PINACONE
	docs, err := utils.GetContentPDF(fileHeader)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ResponseMsg(
			"error", "Error al obtener contenido del archivo",
		))
	}

	texts := []string{}
	for _, d := range docs {
		texts = append(texts, d.PageContent)
	}
	embeddings, err := utils.GetEmbddingsPDF(texts)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ResponseMsg(
			"error", "Error al obtener embeddings del archivo",
		))
	}

	for i := 0; i < len(docs); i += 100 {
		end := i + 100
		if end > len(docs) {
			end = len(docs)
		}
		utils.SaveVectorOnPinacone(docs[i:end], embeddings[i:end], fileKey)
	}

	return c.JSON(utils.ResponseMsg(
		"success", "Chat creado correctamente",
	))

}

func GetAllChats(c *fiber.Ctx) error {
	user := c.Locals("user").(*users.User)
	var chats []Chatpdf
	res := libs.DBInit().
		Order("created_at desc").
		Preload("User").
		Where("user_id = ?", user.ID).
		Find(&chats)
	if res.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(utils.ResponseMsg(
			"error", "Chats no encontrados",
		))
	}
	return c.JSON(utils.Response(
		"success", "Chats encontrados",
		chats,
		fiber.Map{"url": getUrl()},
	))

}

func GetChatsById(c *fiber.Ctx) error {
	id := c.Params("id")
	user := c.Locals("user").(*users.User)
	var chats []Chatpdf

	if err := libs.DBInit().Where("id = ? AND user_id = ?", id, user.ID).Find(&chats).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(utils.ResponseMsg(
			"error", "Chats no encontrados",
		))
	}

	return c.JSON(utils.Response(
		"success", "Chats encontrados",
		chats,
		fiber.Map{"url": getUrl()},
	))

}

func DeleteChat(c *fiber.Ctx) error {
	id := c.Params("id")
	user := c.Locals("user").(*users.User)
	var chat Chatpdf
	if err := libs.DBInit().Where("id = ? AND user_id = ?", id, user.ID).First(&chat).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(utils.ResponseMsg(
			"error", "Chat no encontrado",
		))
	}
	if err := utils.DeleteObject(os.Getenv("AWS_S3_BUCKET"), chat.Key); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ResponseMsg(
			"error", "Error al eliminar archivo",
		))
	}

	if err := utils.DeleteVectorsByNamespace(chat.Key); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ResponseMsg(
			"error", "Error al eliminar vectores",
		))
	}

	if err := libs.DBInit().Select("Messages").Delete(&chat).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ResponseMsg(
			"error", "Error al eliminar chat",
		))
	}

	return c.JSON(utils.ResponseMsg(
		"success", "Chat eliminado correctamente",
	))
}

func GetChatFile(c *fiber.Ctx) error {
	chatId := c.Params("id")
	user := c.Locals("user").(*users.User)
	var chat Chatpdf
	if err := libs.DBInit().Where("id = ? AND user_id = ?", chatId, user.ID).First(&chat).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(utils.ResponseMsg(
			"error", "Chat no encontrado",
		))
	}

	object, err := utils.GetObject(os.Getenv("AWS_S3_BUCKET"), chat.Key)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ResponseMsg(
			"error", "Error al obtener archivo",
		))
	}
	fileBytes, err := io.ReadAll(object.Body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ResponseMsg(
			"error", "Error al obtener archivo",
		))
	}
	fmt.Println(chat.ContentType)
	c.Response().Header.Set("Content-Type", chat.ContentType)
	return c.Send(fileBytes)

}

func GetSimilarity(c *fiber.Ctx) error {
	type Body struct {
		Query string `json:"query"`
	}
	var body Body
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ResponseMsg(
			"error", "Datos incorrectos",
		))
	}

	chatID := c.Params("id")
	user := c.Locals("user").(*users.User)
	var chat Chatpdf
	if err := libs.DBInit().Where("id = ? AND user_id = ?", chatID, user.ID).First(&chat).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(utils.ResponseMsg(
			"error", "Chat no encontrado",
		))
	}

	cxt, err := utils.GetContext(body.Query, chat.Key, 10)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ResponseMsg(
			"error", "Error al obtener contexto",
		))
	}
	return c.JSON(utils.Response(
		"success", "Context fetched successfully",
		cxt,
		nil,
	))
}
