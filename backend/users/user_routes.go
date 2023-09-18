package users

import (
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.RouterGroup) {
	r.POST("/", CreateUser)
	r.PUT("/", UpdateUser)
	r.DELETE("/:social_account", DeleUser)
}
