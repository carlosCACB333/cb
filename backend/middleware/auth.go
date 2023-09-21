package middleware

import (
	"cb/common"
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

func validateClerkToken(token string) *users.User {

	client := libs.ClerkClient()
	sessClaims, err := client.VerifyToken(token)

	if err != nil {
		return nil
	}
	ClerkUser, e := client.Users().Read(sessClaims.Claims.Subject)

	if e != nil {
		return nil
	}
	if ClerkUser == nil {
		return nil
	}
	user := users.User{
		Model: common.Model{
			ID: ClerkUser.ID,
		},
		Status: "active",
	}
	if len(ClerkUser.EmailAddresses) > 0 {
		user.Email = ClerkUser.EmailAddresses[0].EmailAddress
	}
	if len(ClerkUser.PhoneNumbers) > 0 {
		user.Phone = ClerkUser.PhoneNumbers[0].PhoneNumber
	}
	if ClerkUser.Username != nil {
		user.Username = *ClerkUser.Username
	}
	if ClerkUser.FirstName != nil {
		user.FirstName = *ClerkUser.FirstName
	}
	if ClerkUser.LastName != nil {
		user.LastName = *ClerkUser.LastName
	}
	if ClerkUser.ImageURL != nil {
		user.Photo = *ClerkUser.ImageURL
	}
	if ClerkUser.Gender != nil {
		user.Gender = *ClerkUser.Gender
	}

	return &user
}

func validateMyToken(token string) *users.User {
	id, err := libs.ValidateToken(token)
	if err != nil {
		return nil
	}
	var user users.User
	if err := libs.DBInit().Where("id=?", id).First(&user).Error; err != nil {
		return nil
	}
	return &user

}
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionToken := c.Request.Header.Get("Authorization")
		sessionToken = strings.TrimPrefix(sessionToken, "Bearer ")
		ClerkUser := validateClerkToken(sessionToken)
		if ClerkUser == nil {
			myUser := validateMyToken(sessionToken)
			if myUser == nil {
				c.JSON(401, utils.Response(
					"error",
					"Unauthorized",
					nil,
					nil,
				))
				c.Abort()
				return
			}
			c.Set("user", myUser)
			c.Next()
			return
		}

		c.Set("user", ClerkUser)
		c.Next()

	}
}
