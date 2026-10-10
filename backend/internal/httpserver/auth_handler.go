package httpserver

import (
    "github.com/elialombardi/piccis-meditation/backend/pkg/dto"
    "github.com/elialombardi/piccis-meditation/backend/pkg/service"
    "github.com/gofiber/fiber/v3"
)

type AuthHandler struct{ svc *service.AuthService }

func NewAuthHandler(s *service.AuthService) *AuthHandler { return &AuthHandler{svc: s} }

func (h *AuthHandler) Register(c fiber.Ctx) error {
    var req dto.RegisterRequest
    if err := c.Bind().Body(&req); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid body"})
    }
    res, err := h.svc.Register(req)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
    }
    return c.Status(fiber.StatusCreated).JSON(res)
}

func (h *AuthHandler) Login(c fiber.Ctx) error {
    var req dto.LoginRequest
    if err := c.Bind().Body(&req); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid body"})
    }
    res, err := h.svc.Login(req)
    if err != nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
    }
    return c.JSON(res)
}