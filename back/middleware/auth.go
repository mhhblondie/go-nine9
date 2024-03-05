package middleware

import (
	"os"

	"github.com/Go-nine9/go-nine9/helper"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var SecretKey = os.Getenv("JWT_SECRET")

func AuthRequired() fiber.Handler {
	return func(c *fiber.Ctx) error {
		token, err := helper.GetToken(c)
		if token == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized: Token missing",
			})
		}

		claims, err := helper.ParseJWT(token)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized: Invalid Token",
			})
		}

		c.Locals("userClaims", claims)
		return c.Next()
	}
}

func RoleMiddleware(roles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userClaims, ok := c.Locals("userClaims").(jwt.MapClaims)
		if !ok {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Internal Server Error",
			})
		}
		userRole, ok := userClaims["role"].(string)
		if !ok {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"message": "Forbidden: Insufficient Permissions",
			})
		}

		for _, role := range roles {
			if role == userRole {
				return c.Next()
			}
		}

		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "Forbidden: Insufficient Permissions",
		})
	}
}
