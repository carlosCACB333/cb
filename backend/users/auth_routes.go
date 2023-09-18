package users

import (
	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.RouterGroup) {
	r.POST("/register", AuthRegister)
	r.POST("/login", AuthLogin)
	r.POST("/sync", ClerkSync)
	// r.POST("/change-password", middleware.AuthMiddleware(), ChangePassword)
}
