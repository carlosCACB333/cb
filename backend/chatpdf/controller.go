package chatpdf

import (
	"cb/libs"
	"cb/users"

	"cb/utils"
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
)

func getUrl() string {
	return "https://" + os.Getenv("AWS_S3_BUCKET") + ".s3.amazonaws.com/" + "https://" + os.Getenv("AWS_S3_BUCKET") + ".s3.amazonaws.com/"

}
func CreateChatdf(c *gin.Context) {
	c.Request.FormFile("file")
	fileHeader, err := c.FormFile("file")
	if err != nil {
		c.JSON(400, utils.Response(
			"error", "Invalid file",
			nil,
			nil,
		))
		return
	}

	fileName := fileHeader.Filename
	fileKey := time.Now().Format(time.RFC3339) + "-" + fileName

	if fileHeader.Size > 1025*1024*10 {
		c.JSON(400, utils.Response(
			"error", "File size too large",
			nil,
			nil,
		))
		return
	}
	fmt.Println(fileHeader.Header.Get("Content-Type"))

	er := utils.PutObject(fileHeader, os.Getenv("AWS_S3_BUCKET"), fileKey)
	if er != nil {
		c.JSON(400, utils.Response(
			"error", "Unable to upload file",
			nil,
			nil,
		))

		return
	}
	user := c.MustGet("user").(users.User)
	chat := Chatpdf{
		Name:   fileName,
		Key:    fileKey,
		UserID: user.ID,
	}
	res := libs.DBInit().Create(&chat)
	if res.Error != nil {
		c.JSON(400, utils.Response(
			"error", "Unable to create chat",
			nil,
			nil,
		))
		return
	}

	c.JSON(200, utils.Response(
		"success", "Chat created successfully",
		chat,
		gin.H{"url": getUrl()},
	))

}

func GetChatpdfs(c *gin.Context) {
	var chats []Chatpdf
	res := libs.DBInit().Preload("User").Find(&chats)
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

func GetChatpdf(c *gin.Context) {
	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "Hello!",
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}

	fmt.Println(resp.Choices[0].Message.Content)
}
func GetChatpdf3(c *gin.Context) {
	id := c.Param("id")
	var chat Chatpdf
	res := libs.DBInit().Preload("User").Where("id = ?", id).First(&chat)
	if res.Error != nil {
		c.JSON(400, utils.Response(
			"error", "Unable to fetch chat",
			nil,
			nil,
		))
		return
	}

	s3Res, err := utils.GetObject(os.Getenv("AWS_S3_BUCKET"), chat.Key)
	if err != nil {
		c.JSON(400, utils.Response(
			"error", "Unable to get file",
			nil,
			nil,
		))
		return
	}
	doc, err := utils.GetContentPDF(s3Res.Body)
	if err != nil {
		c.JSON(400, utils.Response(
			"error", "Unable to process file",
			nil,
			nil,
		))
		return
	}
	utils.ProcessDocPage(doc[0])

	c.JSON(200, utils.Response(
		"success", "Chat fetched successfully",
		doc,
		nil,
	))

}
func GetChatpdfFile(c *gin.Context) {
	key := c.Param("key")

	object, err := utils.GetObject(os.Getenv("AWS_S3_BUCKET"), key)
	if err != nil {
		c.JSON(400, utils.Response(
			"error", "Unable to fetch chat",
			nil,
			nil,
		))
		return
	}
	fileBytes, err := io.ReadAll(object.Body)
	if err != nil {
		c.JSON(400, utils.Response(
			"error", "Unable to fetch chat",
			nil,
			nil,
		))
		return
	}

	c.Data(200, object.Metadata["Content-Type"], fileBytes)

}
