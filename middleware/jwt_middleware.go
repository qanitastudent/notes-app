package middleware

import (
	"os"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

// JWTProtected mengembalikan middleware Fiber JWT
func JWTProtected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
		ContextKey: "user", // untuk akses user di handler
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
		},
	})
}

// Ambil user_id dari JWT di c.Locals("user")
func GetUserID(c *fiber.Ctx) int {
	userToken := c.Locals("user")
	if userToken == nil {
		return 0
	}

	token := userToken.(*jwtware.Token)
	claims := token.Claims.(jwtware.MapClaims)
	if uid, ok := claims["user_id"].(float64); ok {
		return int(uid)
	}
	return 0
}
