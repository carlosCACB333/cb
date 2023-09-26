package chatpdf

import (
	"github.com/gin-gonic/gin"
)

func ChatPdfRoutes(r *gin.RouterGroup) {
	r.GET("", GetAllChats)
	r.GET("/:id", GetChatsById)
	r.DELETE("/:id", DeleteChat)
}
func BootRoutesRoutes(r *gin.RouterGroup) {
	r.POST("/:id", GetSimilarity)
}

func MessagesRoutes(r *gin.RouterGroup) {
	r.POST("/:id/new", CreateMessage)
	r.GET("/:id/last", GetLastMessages)
}
