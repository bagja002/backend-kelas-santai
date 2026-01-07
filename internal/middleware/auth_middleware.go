package middleware

import (
	"project-kelas-santai/pkg/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Protected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return utils.SendError(c, fiber.StatusUnauthorized, "Missing Authorization header", nil)
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
		token, err := utils.ValidateToken(tokenString)
		if err != nil {
			return utils.SendError(c, fiber.StatusUnauthorized, "Invalid or expired token", err.Error())
		}

		// Store claims in context for valid token
		claims := token.Claims.(jwt.MapClaims)
		c.Locals("user_id", claims["user_id"])
		c.Locals("role", claims["role"])

		return c.Next()
	}
}

func AdminProtected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		role := c.Locals("role")
		if role != "admin" {
			return utils.SendError(c, fiber.StatusForbidden, "Access denied: Admins only", nil)
		}
		return c.Next()
	}
}
