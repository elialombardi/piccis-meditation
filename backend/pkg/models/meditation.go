package models

import (
    "encoding/json"
    "time"
)

// MeditationType is a catalog entry (e.g. "Breathwork", "Body Scan").
// It is global (not user-scoped) and seeded on migration.
type MeditationType struct {
    ID          uint            `gorm:"primaryKey" json:"id"`
    Slug        string          `gorm:"uniqueIndex;size:64;not null" json:"slug"`
    Name        string          `gorm:"size:128;not null" json:"name"`
    Description string          `gorm:"type:text" json:"description"`
    Icon        string          `gorm:"size:64" json:"icon"`     // MUI icon name
    Script      json.RawMessage `gorm:"type:jsonb" json:"script"` // phases: see below
    CreatedAt   time.Time       `json:"createdAt"`
}

// Script JSON shape (stored on MeditationType.Script):
// {
//   "phases": [
//     {"label":"Inhale","seconds":4,"speak":"Breathe in slowly"},
//     {"label":"Hold","seconds":4,"speak":"Hold"},
//     {"label":"Exhale","seconds":6,"speak":"Release"}
//   ],
//   "repeat": 8
// }

// MeditationSession is one completed (or in-progress) sit by a user.
type MeditationSession struct {
    ID               uint      `gorm:"primaryKey" json:"id"`
    UserID           uint      `gorm:"index;not null" json:"userId"`
    MeditationTypeID uint      `gorm:"index;not null" json:"meditationTypeId"`
    DurationSeconds  int       `gorm:"not null" json:"durationSeconds"`
    CompletedAt      time.Time `gorm:"index" json:"completedAt"`
    Notes            string    `gorm:"type:text" json:"notes"`
    CreatedAt        time.Time `json:"createdAt"`
}