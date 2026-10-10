package httpserver

import (
	"log"

	"github.com/elialombardi/piccis-meditation/backend/internal/config"
	"github.com/elialombardi/piccis-meditation/backend/pkg/service"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

// AppDeps wires the handlers and shared config the routes need.
type AppDeps struct {
	AuthHandler       *AuthHandler
	MeditationHandler *MeditationHandler
	JWTSecret         string
}

// New builds the Fiber app. Signature: New(cfg, authSvc, medSvc).
func New(cfg *config.Config, authSvc *service.AuthService, medSvc *service.MeditationService) *fiber.App {
	app := fiber.New(fiber.Config{AppName: "Picci's Meditation"})

	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization"},
	}))

	app.Get("/health", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	RegisterRoutes(app, &AppDeps{
		AuthHandler:       NewAuthHandler(authSvc),
		MeditationHandler: NewMeditationHandler(medSvc),
		JWTSecret:         cfg.JWTSecret,
	})

	log.Printf("listening on %s", cfg.HTTPAddr)
	return app
}