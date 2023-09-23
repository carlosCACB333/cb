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

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func getUrl() string {
	return "https://" + os.Getenv("AWS_S3_BUCKET") + ".s3.amazonaws.com/"
}

func CreateChat(c *gin.Context) {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		fmt.Println(err)
		c.JSON(400, utils.Response(
			"error", "Invalid file",
			nil,
			nil,
		))
		return
	}

	if fileHeader.Size > 1025*1024*10 {
		c.JSON(400, utils.Response(
			"error", "File size too large",
			nil,
			nil,
		))
		return
	}
	fileName := fileHeader.Filename
	ID := uuid.New().String()
	split := strings.Split(fileName, ".")
	fileKey := ID + "." + split[len(split)-1]

	er := utils.PutObject(fileHeader, os.Getenv("AWS_S3_BUCKET"), fileKey)
	if er != nil {
		c.JSON(400, utils.Response(
			"error", "Unable to upload file",
			nil,
			nil,
		))

		return
	}
	user := c.MustGet("user").(*users.User)

	chat := Chatpdf{
		Model:       common.Model{ID: ID},
		Name:        fileName,
		Key:         fileKey,
		UserID:      user.ID,
		ContentType: fileHeader.Header.Get("Content-Type"),
	}

	if err := libs.DBInit().Create(&chat).Error; err != nil {
		c.JSON(400, utils.Response(
			"error", "Unable to create chat",
			nil,
			nil,
		))
		return
	}
	// VECTORIZE AND SAVE ON PINACONE
	docs, err := utils.GetContentPDF(fileHeader)
	if err != nil {
		c.JSON(400, utils.Response(
			"error", "Unable to process file",
			nil,
			nil,
		))
		return
	}

	texts := []string{}
	for _, d := range docs {
		texts = append(texts, d.PageContent)
	}
	embeddings, err := utils.GetEmbddingsPDF(texts)
	if err != nil {
		c.JSON(400, utils.Response(
			"error", "Unable to create encodings",
			nil,
			nil,
		))
		return
	}

	for i := 0; i < len(docs); i += 100 {
		end := i + 100
		if end > len(docs) {
			end = len(docs)
		}
		utils.SaveVectorOnPinacone(docs[i:end], embeddings[i:end], fileKey)
	}

	c.JSON(200, utils.Response(
		"success", "Chat created successfully",
		chat,
		gin.H{"url": getUrl()},
	))

}

func GetAllChats(c *gin.Context) {
	user := c.MustGet("user").(*users.User)
	var chats []Chatpdf
	res := libs.DBInit().
		Order("created_at desc").
		Preload("User").
		Where("user_id = ?", user.ID).
		Find(&chats)
	if res.Error != nil {
		c.JSON(400, utils.Response(
			"error", "Unable to fetch chats",
			nil,
			nil,
		))
		return
	}
	c.JSON(200, utils.Response(
		"success", "Chats fetched successfully",
		chats,
		nil,
	))

}

func GetMessagesByChatId(c *gin.Context) {
	id := c.Param("id")
	user := c.MustGet("user").(*users.User)
	var chats []Chatpdf

	if err := libs.DBInit().Where("id = ? AND user_id = ?", id, user.ID).Find(&chats).Error; err != nil {
		c.JSON(400, utils.Response(
			"error", "Unable to fetch chat",
			nil,
			nil,
		))
		return
	}

	c.JSON(200, utils.Response(
		"success", "Chat fetched successfully",
		chats,
		gin.H{"url": getUrl()},
	))

}
func GetChatFile(c *gin.Context) {
	chatId := c.Param("id")
	user := c.MustGet("user").(*users.User)
	var chat Chatpdf
	if err := libs.DBInit().Where("id = ? AND user_id = ?", chatId, user.ID).First(&chat).Error; err != nil {
		c.JSON(404, utils.ResponseMsg("Resource not found"))
		return
	}

	object, err := utils.GetObject(os.Getenv("AWS_S3_BUCKET"), chat.Key)
	if err != nil {
		c.JSON(404, utils.ResponseMsg("Resource not found"))
		return
	}
	fileBytes, err := io.ReadAll(object.Body)
	if err != nil {
		c.JSON(400, utils.ResponseMsg("Unable to fetch file"))
		return
	}
	fmt.Println(object.Metadata["Content-Type"])
	c.Data(200, chat.ContentType, fileBytes)

}

func GetSimilarity(c *gin.Context) {
	type Body struct {
		Query string `json:"query"`
	}
	var body Body
	c.BindJSON(&body)
	chatID := c.Param("id")
	user := c.MustGet("user").(*users.User)
	var chat Chatpdf
	if err := libs.DBInit().Where("id = ? AND user_id = ?", chatID, user.ID).First(&chat).Error; err != nil {
		c.JSON(400, utils.ResponseMsg("Unable to fetch chat"))
		return
	}

	cxt, err := utils.GetContext(body.Query, chat.Key, 10)
	if err != nil {
		c.JSON(400, utils.ResponseMsg("Unable to fetch context"))
		return
	}
	c.JSON(200, utils.Response(
		"success", "Context fetched successfully",
		cxt,
		nil,
	))
}
