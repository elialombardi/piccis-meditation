package httpserver

import (
    "github.com/elialombardi/piccis-meditation/backend/internal/httpserver/middleware"
    "github.com/gofiber/fiber/v3"
)

func RegisterRoutes(app *fiber.App, cfg *AppDeps) {
    api := app.Group("/api")

    auth := api.Group("/auth")
    auth.Post("/register", cfg.AuthHandler.Register)
    auth.Post("/login", cfg.AuthHandler.Login)

    // Public catalog
    api.Get("/meditation/types", cfg.MeditationHandler.ListTypes)

    // User-scoped
    med := api.Group("/meditation", middleware.JWT(cfg.JWTSecret))
    med.Post("/sessions", cfg.MeditationHandler.CreateSession)
    med.Get("/sessions", cfg.MeditationHandler.ListSessions)
    med.Get("/stats", cfg.MeditationHandler.Stats)
}