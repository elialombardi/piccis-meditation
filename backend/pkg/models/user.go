package models

import "time"

type User struct {
    ID           uint      `gorm:"primaryKey" json:"id"`
    Email        string    `gorm:"uniqueIndex;size:255;not null" json:"email"`
    PasswordHash string    `gorm:"not null" json:"-"`
    CreatedAt    time.Time `json:"createdAt"`
    UpdatedAt    time.Time `json:"updatedAt"`
}