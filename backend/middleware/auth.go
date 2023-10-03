package middleware

import (
	"cb/common"
	"cb/libs"
	"cb/users"
	"cb/utils"
	"fmt"

	"strings"

	"github.com/gofiber/fiber/v2"
)

func ApiKeyMiddleware(apiKey string) fiber.Handler {

	return func(c *fiber.Ctx) error {
		xApiKey := c.Request().Header.Peek("x-api-key")
		if string(xApiKey) != apiKey {
			return c.Status(fiber.StatusUnauthorized).JSON(utils.ResponseMsg(
				"error",
				"Unauthorized",
			))

		}
		return c.Next()
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
func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		fmt.Println("AuthMiddleware")
		auth := c.Request().Header.Peek("Authorization")
		sessionToken := (strings.TrimPrefix(string(auth), "Bearer "))
		ClerkUser := validateClerkToken(sessionToken)
		if ClerkUser == nil {
			myUser := validateMyToken(sessionToken)
			if myUser == nil {
				return c.Status(fiber.StatusUnauthorized).JSON(utils.ResponseMsg(
					"error",
					"Unauthorized",
				))
			}
			c.Locals("user", myUser)
			return c.Next()
		}
		c.Locals("user", ClerkUser)
		return c.Next()

	}
}
