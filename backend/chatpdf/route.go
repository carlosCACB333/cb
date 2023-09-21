package chatpdf

import (
	"github.com/gin-gonic/gin"
)

func ChatPdfRoutes(r *gin.RouterGroup) {
	r.POST("", CreateChat)
	r.GET("", GetAllChats)
	r.GET("/:id", GetMessagesByChatId)
	r.GET("/resource/:key", GetChatFile)
}
func BootRoutesRoutes(r *gin.RouterGroup) {
	r.POST("/:id", GetSimilarity)
}

func MessagesRoutes(r *gin.RouterGroup) {
	r.POST("/:id/new", CreateMessage)
	r.GET("/:id/last", GetLastMessages)
}
