package chatpdf

import (
	"cb/common"
	"cb/libs"
	"cb/users"
	"fmt"
	"strings"
	"sync"
	"time"

	"cb/utils"
	"io"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sashabaranov/go-openai"
	"github.com/tmc/langchaingo/schema"
)

func getUrl() string {
	return "https://" + os.Getenv("AWS_S3_BUCKET") + ".s3.amazonaws.com/"
}

func CreateChat(c *fiber.Ctx) error {
	start := time.Now()
	fileHeader, err := c.FormFile("file")
	if err != nil {
		fmt.Println(err)
		return utils.NewError(fiber.StatusBadRequest, "Error al obtener archivo", nil)
	}

	if fileHeader.Size > 1025*1024*10 {
		return utils.NewError(fiber.StatusBadRequest, "El archivo es demasiado grande", nil)
	}

	fileName := fileHeader.Filename
	ID := uuid.New().String()
	split := strings.Split(fileName, ".")
	fileKey := ID + "." + split[len(split)-1]
	user := c.Locals("user").(*users.User)

	// SAVE FILE ON S3
	var uploadChain = make(chan error)
	defer close(uploadChain)
	go func() {
		uploadChain <- utils.PutObject(fileHeader, os.Getenv("AWS_S3_BUCKET"), fileKey)
	}()

	// VECTORIZE FILE
	var vectorizeChain = make(chan []schema.Document)
	defer close(vectorizeChain)
	go func() {
		vectorizeChain <- utils.GetContentPDF(fileHeader)
	}()

	// WAIT FOR UPLOAD AND VECTORIZE
	var docs []schema.Document
	for i := 0; i < 2; i++ {
		select {
		case err := <-uploadChain:
			if err != nil {
				return utils.NewError(fiber.StatusBadRequest, "Error al subir archivo", nil)
			}
		case d := <-vectorizeChain:
			if d == nil {
				return utils.NewError(fiber.StatusBadRequest, "Error al obtener contenido del archivo", nil)
			}
			docs = d
		}
	}

	texts := []string{}
	for _, d := range docs {
		texts = append(texts, d.PageContent)
	}
	embeddings, err := utils.GetEmbddingsPDF(texts)
	if err != nil {
		return utils.NewError(fiber.StatusBadRequest,
			"No se pudo obtener embeddings del archivo", nil)
	}

	// SAVE VECTORS ON PINACONE
	var wg sync.WaitGroup
	for i := 0; i < len(docs); i += 100 {
		wg.Add(1)
		end := i + 100
		if end > len(docs) {
			end = len(docs)
		}
		go func(d []schema.Document, e []openai.Embedding) {
			defer wg.Done()
			utils.SaveVectorOnPinacone(d, e, fileKey)

		}(docs[i:end], embeddings[i:end])

	}

	chat := Chatpdf{
		Model:       common.Model{ID: ID},
		Name:        fileName,
		Key:         fileKey,
		UserID:      user.ID,
		ContentType: fileHeader.Header.Get("Content-Type"),
	}

	resultInsert := libs.DBInit().Create(&chat)
	wg.Wait()

	if resultInsert.Error != nil {
		return utils.NewError(fiber.StatusBadRequest, "Error al crear chat", nil)
	}

	fmt.Println("Time: ", time.Since(start))

	return c.JSON(utils.NewBody(utils.Body{
		Data:    chat,
		Others:  fiber.Map{"url": getUrl()},
		Message: "Chat creado correctamente",
	}))

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
		return utils.Error{
			Code:    fiber.StatusInternalServerError,
			Message: "Error al obtener chats",
		}
	}
	return c.JSON(utils.NewBody(utils.Body{
		Data:   chats,
		Others: fiber.Map{"url": getUrl()},
	}))

}

func GetChatsById(c *fiber.Ctx) error {
	id := c.Params("id")
	user := c.Locals("user").(*users.User)
	var chats []Chatpdf

	if err := libs.DBInit().Where("id = ? AND user_id = ?", id, user.ID).Find(&chats).Error; err != nil {
		return utils.NewError(fiber.StatusNotFound, "Chat no encontrado", nil)
	}

	return c.JSON(utils.NewBody(utils.Body{
		Data:   chats,
		Others: fiber.Map{"url": getUrl()},
	}))

}

func DeleteChat(c *fiber.Ctx) error {
	start := time.Now()
	id := c.Params("id")
	user := c.Locals("user").(*users.User)
	var chat Chatpdf
	if err := libs.DBInit().Where("id = ? AND user_id = ?", id, user.ID).First(&chat).Error; err != nil {
		return utils.NewError(fiber.StatusNotFound, "Chat no encontrado", nil)

	}
	// Delete concurrente
	var s3Chain = make(chan error)
	var piconeChain = make(chan error)
	var dbChain = make(chan error)
	defer close(s3Chain)
	defer close(piconeChain)
	defer close(dbChain)
	go func() {
		s3Chain <- utils.DeleteObject(os.Getenv("AWS_S3_BUCKET"), chat.Key)
	}()
	go func() {
		piconeChain <- utils.DeleteVectorsByNamespace(chat.Key)
	}()
	go func() {
		dbChain <- libs.DBInit().Select("Messages").Delete(&chat).Error
	}()

	for i := 0; i < 3; i++ {
		select {
		case err := <-s3Chain:
			if err != nil {
				return utils.NewError(fiber.StatusBadRequest, "Error al eliminar archivo", nil)
			}
		case err := <-piconeChain:
			if err != nil {
				return utils.NewError(fiber.StatusBadRequest, "Error al eliminar vectores", nil)
			}
		case err := <-dbChain:
			if err != nil {
				return utils.NewError(fiber.StatusBadRequest, "Error al eliminar chat", nil)
			}
		}
	}

	fmt.Println("Time: ", time.Since(start))

	return c.JSON(utils.NewBody(utils.Body{
		Message: "Chat eliminado correctamente",
	}))

}

func GetChatFile(c *fiber.Ctx) error {
	chatId := c.Params("id")
	user := c.Locals("user").(*users.User)
	var chat Chatpdf
	if err := libs.DBInit().Where("id = ? AND user_id = ?", chatId, user.ID).First(&chat).Error; err != nil {
		return utils.NewError(fiber.StatusNotFound, "Chat no encontrado", nil)

	}

	object, err := utils.GetObject(os.Getenv("AWS_S3_BUCKET"), chat.Key)
	if err != nil {
		return utils.NewError(fiber.StatusBadRequest, "Error al obtener archivo", nil)
	}
	fileBytes, err := io.ReadAll(object.Body)
	if err != nil {
		return utils.NewError(fiber.StatusBadRequest, "Error al obtener archivo", nil)
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
		return utils.NewError(fiber.StatusBadRequest, "Datos incorrectos", nil)
	}

	chatID := c.Params("id")
	user := c.Locals("user").(*users.User)
	var chat Chatpdf
	if err := libs.DBInit().Where("id = ? AND user_id = ?", chatID, user.ID).First(&chat).Error; err != nil {
		return utils.NewError(fiber.StatusNotFound, "Chat no encontrado", nil)
	}

	cxt, err := utils.GetContext(body.Query, chat.Key, 10)
	if err != nil {
		return utils.NewError(fiber.StatusBadRequest, "Error al obtener contexto", nil)
	}
	return c.JSON(utils.NewBody(utils.Body{
		Data: cxt,
	}))

}
