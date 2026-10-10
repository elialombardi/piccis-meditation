package middleware

import (
    "strings"

    "github.com/elialombardi/piccis-meditation/backend/pkg/auth"
    "github.com/gofiber/fiber/v3"
)

func JWT(secret string) fiber.Handler {
    return func(c fiber.Ctx) error {
        h := c.Get("Authorization")
        if !strings.HasPrefix(h, "Bearer ") {
            return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "missing token"})
        }
        claims, err := auth.Parse(secret, strings.TrimPrefix(h, "Bearer "))
        if err != nil {
            return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid token"})
        }
        c.Locals("userID", claims.UserID)
        return c.Next()
    }
}