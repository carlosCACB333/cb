package middleware

import (
	"cb/lib"
	"cb/model"
	"cb/server"
	"cb/util"

	"strings"

	"github.com/gofiber/fiber/v2"
)

func ApiKeyMiddleware(apiKey string) fiber.Handler {

	return func(c *fiber.Ctx) error {
		xApiKey := c.Request().Header.Peek("x-api-key")
		if string(xApiKey) != apiKey {
			return util.NewError(fiber.StatusUnauthorized, "Unauthorized", nil)
		}
		return c.Next()
	}
}

func validateClerkToken(token string) string {

	client := lib.ClerkClient()
	sessClaims, err := client.VerifyToken(token)

	if err != nil {
		return ""
	}
	ClerkUser, e := client.Users().Read(sessClaims.Claims.Subject)

	if e != nil || ClerkUser == nil {
		return ""
	}

	return ClerkUser.ID
}

func validateMyToken(token string) string {
	id, err := lib.ValidateToken(token)
	if err != nil {
		return ""
	}
	return id
}
func AuthMiddleware(s *server.Server) fiber.Handler {
	return func(c *fiber.Ctx) error {
		auth := c.Request().Header.Peek("Authorization")
		sessionToken := (strings.TrimPrefix(string(auth), "Bearer "))
		resultChain := make(chan string)

		go func() {
			resultChain <- validateClerkToken(sessionToken)
		}()
		go func() {
			resultChain <- validateMyToken(sessionToken)
		}()

		var userID string
		for i := 0; i < 2; i++ {
			if userID = <-resultChain; userID != "" {
				break
			}
		}

		if userID == "" {
			return util.NewError(fiber.StatusUnauthorized, "Unauthorized", nil)
		}

		var user model.User
		if err := s.DB().Where("id = ?", userID).First(&user).Error; err != nil {
			return util.NewError(fiber.StatusUnauthorized, "Unauthorized", nil)
		}

		c.Locals("user", &user)
		return c.Next()
	}
}
