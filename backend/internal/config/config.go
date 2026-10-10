package config

import (
    "os"
    "strconv"
    "time"

    "github.com/joho/godotenv"
)

type Config struct {
    Env         string
    HTTPAddr    string
    DatabaseURL string
    JWTSecret   string
    JWTTTL      time.Duration
}

func Load() *Config {
    _ = godotenv.Load()
    ttl, _ := strconv.Atoi(getEnv("JWT_TTL_HOURS", "24"))
    return &Config{
        Env:         getEnv("APP_ENV", "development"),
        HTTPAddr:    getEnv("HTTP_ADDR", ":8080"),
        DatabaseURL: getEnv("DATABASE_URL", ""),
        JWTSecret:   getEnv("JWT_SECRET", "dev-secret"),
        JWTTTL:      time.Duration(ttl) * time.Hour,
    }
}

func getEnv(k, def string) string {
    if v := os.Getenv(k); v != "" {
        return v
    }
    return def
}