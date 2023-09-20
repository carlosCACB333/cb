package middleware

import (
	"cb/libs"
	"cb/utils"
	"fmt"

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
		sessionToken := c.Request.Header.Get("Authorization")
		sessionToken = strings.TrimPrefix(sessionToken, "Bearer ")
		fmt.Println(sessionToken)
		client := libs.ClerkClient()
		sessClaims, err := client.VerifyToken(sessionToken)

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
		ClerkUser, err := client.Users().Read(sessClaims.Claims.Subject)
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
		c.Set("user", ClerkUser)
		c.Next()

	}
}
