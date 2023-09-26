package chatpdf

import (
	"cb/libs"
	"cb/users"
	"cb/utils"
	"sort"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateMessage(c *gin.Context) {
	chatId := c.Param("id")
	chat := Chatpdf{}
	db := libs.DBInit()
	user := c.MustGet("user").(*users.User)
	if err := db.Where("id = ? AND user_id = ?", chatId, user.ID).First(&chat).Error; err != nil {
		c.JSON(404, utils.ResponseMsg("error", "Chat not found"))
		return
	}

	var message ChatpdfMessage
	if err := c.ShouldBindJSON(&message); err != nil {
		c.JSON(400, utils.ResponseMsg("error", "Invalid request"))
		return
	}
	message.ChatpdfID = chat.ID
	message.ID = uuid.New().String()
	if err := utils.ValidateFields(message); err != nil {
		c.JSON(400, utils.Response(
			"error", "Data validation error",
			err,
			nil,
		))
		return
	}
	if err := db.Create(&message).Error; err != nil {
		c.JSON(400, utils.ResponseMsg("error", "Unable to create message"))
		return
	}
	c.JSON(200, utils.Response(
		"success", "Message created successfully",
		message,
		nil,
	))

}

func GetLastMessages(c *gin.Context) {
	db := libs.DBInit()
	chatId := c.Param("id")
	user := c.MustGet("user").(*users.User)
	var chat Chatpdf
	var messages []ChatpdfMessage

	// validate if chat exists and belongs to user
	if err := db.Where("id = ? AND user_id = ?", chatId, user.ID).First(&chat).Error; err != nil {
		c.JSON(404, utils.ResponseMsg("error", "Chat not found"))
		return
	}

	if err := db.Where("chatpdf_id = ?", chat.ID).Order("created_at desc").Limit(6).Find(&messages).Error; err != nil {
		c.JSON(400, utils.ResponseMsg("error", "Unable to fetch messages"))
		return
	}

	sort.Slice(messages, func(i, j int) bool {
		return messages[i].CreatedAt.Before(messages[j].CreatedAt)
	})

	c.JSON(200, utils.Response(
		"success", "Messages fetched successfully",
		messages,
		nil,
	))

}
