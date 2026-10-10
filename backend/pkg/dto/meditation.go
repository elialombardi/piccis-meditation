package dto

import "time"

type MeditationTypeDTO struct {
    ID          uint            `json:"id"`
    Slug        string          `json:"slug"`
    Name        string          `json:"name"`
    Description string          `json:"description"`
    Icon        string          `json:"icon"`
    Script      map[string]any  `json:"script"`
}

type CreateSessionRequest struct {
    MeditationTypeID uint   `json:"meditationTypeId" validate:"required"`
    DurationSeconds  int    `json:"durationSeconds" validate:"required,min=1"`
    Notes            string `json:"notes"`
}

type SessionDTO struct {
    ID               uint      `json:"id"`
    MeditationTypeID uint      `json:"meditationTypeId"`
    TypeName         string    `json:"typeName"`
    DurationSeconds  int       `json:"durationSeconds"`
    CompletedAt      time.Time `json:"completedAt"`
    Notes            string    `json:"notes"`
}

type HabitStatsDTO struct {
    TotalSessions     int            `json:"totalSessions"`
    TotalMinutes      int            `json:"totalMinutes"`
    CurrentStreakDays int            `json:"currentStreakDays"`
    LongestStreakDays int            `json:"longestStreakDays"`
    ByType            map[string]int `json:"byType"`         // slug -> count
    Last30Days        []DailyCount   `json:"last30Days"`      // chart data
}

type DailyCount struct {
    Date  string `json:"date"`  // YYYY-MM-DD
    Count int    `json:"count"`
}