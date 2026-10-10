package db

import (
    "log"

    "github.com/elialombardi/piccis-meditation/backend/pkg/models"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

func Connect(dsn string) *gorm.DB {
    g, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("db connect: %v", err)
    }
    if err := g.AutoMigrate(
        &models.User{},
        &models.MeditationType{},
        &models.MeditationSession{},
    ); err != nil {
        log.Fatalf("db migrate: %v", err)
    }
    seedMeditationTypes(g)
    return g
}

// seedMeditationTypes is idempotent: it inserts each type only if the slug is missing.
func seedMeditationTypes(g *gorm.DB) {
    types := []models.MeditationType{
        {
            Slug:        "breathwork",
            Name:        "Breathwork",
            Description: "Box-breathing with a 4-4-4-4 rhythm. Great for quick calm.",
            Icon:        "Air",
            Script: []byte(`{
                "repeat": 8,
                "phases": [
                    {"label":"Inhale","seconds":4,"speak":"Breathe in"},
                    {"label":"Hold","seconds":4,"speak":"Hold"},
                    {"label":"Exhale","seconds":4,"speak":"Breathe out"},
                    {"label":"Rest","seconds":4,"speak":"Rest"}
                ]
            }`),
        },
        {
            Slug:        "body-scan",
            Name:        "Body Scan",
            Description: "Move attention from head to toe, releasing tension.",
            Icon:        "SelfImprovement",
            Script: []byte(`{
                "repeat": 1,
                "phases": [
                    {"label":"Head","seconds":20,"speak":"Notice your head and face. Relax."},
                    {"label":"Shoulders","seconds":20,"speak":"Let your shoulders drop."},
                    {"label":"Chest","seconds":20,"speak":"Feel your chest rise and fall."},
                    {"label":"Hands","seconds":20,"speak":"Relax your hands."},
                    {"label":"Belly","seconds":20,"speak":"Soften your belly."},
                    {"label":"Legs","seconds":20,"speak":"Feel your legs heavy and grounded."},
                    {"label":"Feet","seconds":20,"speak":"Notice your feet."}
                ]
            }`),
        },
        {
            Slug:        "loving-kindness",
            Name:        "Loving-Kindness",
            Description: "Cultivate warmth toward yourself and others.",
            Icon:        "FavoriteBorder",
            Script: []byte(`{
                "repeat": 1,
                "phases": [
                    {"label":"Yourself","seconds":45,"speak":"May I be happy. May I be well."},
                    {"label":"Someone you love","seconds":45,"speak":"May you be happy. May you be well."},
                    {"label":"Someone neutral","seconds":45,"speak":"May we all be happy."},
                    {"label":"All beings","seconds":45,"speak":"May all beings be at peace."}
                ]
            }`),
        },
    }
    for _, t := range types {
        var existing models.MeditationType
        if err := g.Where("slug = ?", t.Slug).First(&existing).Error; err == gorm.ErrRecordNotFound {
            g.Create(&t)
        }
    }
}