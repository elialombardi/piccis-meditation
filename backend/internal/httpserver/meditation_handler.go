package httpserver

import (
    "strconv"

    "github.com/elialombardi/piccis-meditation/backend/pkg/dto"
    "github.com/elialombardi/piccis-meditation/backend/pkg/service"
    "github.com/gofiber/fiber/v3"
)

type MeditationHandler struct{ svc *service.MeditationService }

func NewMeditationHandler(s *service.MeditationService) *MeditationHandler {
    return &MeditationHandler{svc: s}
}

func (h *MeditationHandler) ListTypes(c fiber.Ctx) error {
    types, err := h.svc.ListTypes()
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": err.Error()})
    }
    return c.JSON(types)
}

func (h *MeditationHandler) CreateSession(c fiber.Ctx) error {
    uid := c.Locals("userID").(uint)
    var req dto.CreateSessionRequest
    if err := c.Bind().Body(&req); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "invalid body"})
    }
    sess, err := h.svc.CreateSession(uid, req)
    if err != nil {
        return c.Status(400).JSON(fiber.Map{"error": err.Error()})
    }
    return c.Status(201).JSON(sess)
}

func (h *MeditationHandler) ListSessions(c fiber.Ctx) error {
    uid := c.Locals("userID").(uint)
    limit, _ := strconv.Atoi(c.Query("limit", "50"))
    sessions, err := h.svc.ListSessions(uid, limit)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": err.Error()})
    }
    return c.JSON(sessions)
}

func (h *MeditationHandler) Stats(c fiber.Ctx) error {
    uid := c.Locals("userID").(uint)
    stats, err := h.svc.Stats(uid)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": err.Error()})
    }
    return c.JSON(stats)
}