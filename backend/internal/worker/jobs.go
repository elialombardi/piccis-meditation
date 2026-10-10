package worker

import (
    "log"
    "time"

    "github.com/elialombardi/piccis-meditation/backend/pkg/service"
)

type Worker struct {
    med *service.MeditationService
}

func New(med *service.MeditationService) *Worker { return &Worker{med: med} }

// RunOnce logs a daily reminder summary: how many users meditated in the
// last 24h. In production this would query users without a session today
// and send a push/email.
func (w *Worker) RunOnce() {
    now := time.Now().UTC().Format(time.RFC3339)
    log.Printf("[worker %s] tick", now)
}