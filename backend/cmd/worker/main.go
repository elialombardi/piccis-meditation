package main

import (
    "log"
    "time"

    "github.com/elialombardi/piccis-meditation/backend/internal/config"
    "github.com/elialombardi/piccis-meditation/backend/internal/worker"
    "github.com/elialombardi/piccis-meditation/backend/pkg/db"
    "github.com/elialombardi/piccis-meditation/backend/pkg/service"
)

func main() {
    cfg := config.Load()
    g := db.Connect(cfg.DatabaseURL)
    medSvc := service.NewMeditationService(g)
    w := worker.New(medSvc)

    log.Println("worker started")
    ticker := time.NewTicker(60 * time.Second)
    defer ticker.Stop()
    for range ticker.C {
        w.RunOnce()
    }
}