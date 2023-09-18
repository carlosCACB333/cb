package middleware

import (
	"cb/libs"
	"cb/users"
	"cb/utils"

	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func ApiKeyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		xApiKey := c.Request.Header.Get("X-API-KEY")
		if xApiKey != os.Getenv("X_API_KEY") {
			c.JSON(401, utils.Response(
				"error",
				"Unauthorized",
				nil,
				nil,
			))
			c.Abort()
			return
		}
		c.Next()
	}
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := strings.Replace(c.Request.Header.Get("Authorization"), "Bearer ", "", -1)
		uid, err := libs.ValidateToken(token)
		if err != nil {
			c.JSON(401, utils.Response(
				"error",
				"Unauthorized",
				nil,
				nil,
			))
			c.Abort()
			return
		}

		var user users.User
		if err := libs.DBInit().Where("id = ?", uid).First(&user).Error; err != nil {
			c.JSON(401, utils.Response(
				"error",
				"Unauthorized",
				nil,
				nil,
			))
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Next()

	}
}
