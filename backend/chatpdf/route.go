package chatpdf

import (
	"github.com/gin-gonic/gin"
)

func ChatPdfRoutes(r *gin.RouterGroup) {
	r.POST("", CreateChatdf)
	r.GET("", GetChatpdfs)
	r.GET("/:id", GetChatpdf)
	r.GET("/resource/:key", GetChatpdfFile)
}
