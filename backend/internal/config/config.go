package config

import (
    "github.com/joho/godotenv"
    "log"
    "os"
)

type Config struct {
    DBUrl string
    Port  string
}

func Load() *Config {
    _ = godotenv.Load() // .env is optional in prod
    cfg := &Config{
        DBUrl: os.Getenv("DATABASE_URL"),
        Port:  getEnv("PORT", "8080"),
    }
    return cfg
}

func getEnv(key, fallback string) string {
    if v := os.Getenv(key); v != "" {
        return v
    }
    return fallback
}

func MustLoad() *Config {
    cfg := Load()
    if cfg.DBUrl == "" {
        log.Fatal("DATABASE_URL not set")
    }
    return cfg
}
