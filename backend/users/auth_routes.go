package users

import (
	"cb/middleware"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.RouterGroup) {
	r.POST("/register", AuthRegister)
	r.POST("/login", AuthLogin)
	r.POST("/change-password", middleware.AuthMiddleware(), ChangePassword)
}
