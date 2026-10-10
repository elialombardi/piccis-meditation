package main

import (
	"log"

	"github.com/elialombardi/piccis-meditation/backend/internal/config"
	"github.com/elialombardi/piccis-meditation/backend/internal/httpserver"
	"github.com/elialombardi/piccis-meditation/backend/pkg/db"
	"github.com/elialombardi/piccis-meditation/backend/pkg/service"
)

func main() {
	cfg := config.Load()
	g := db.Connect(cfg.DatabaseURL)

	authSvc := service.NewAuthService(g, cfg.JWTSecret, cfg.JWTTTL)
	medSvc := service.NewMeditationService(g)

	app := httpserver.New(cfg, authSvc, medSvc)
	if err := app.Listen(cfg.HTTPAddr); err != nil {
		log.Fatal(err)
	}
}