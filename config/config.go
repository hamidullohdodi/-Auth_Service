package config

import (
    "log"
    "os"
    "path/filepath"

    "github.com/joho/godotenv"
    "github.com/spf13/cast"
)

type Config struct {
    HTTP_PORT   string
    DB_HOST     string
    DB_PORT     int
    DB_USER     string
    DB_PASSWORD string
    DB_NAME     string
}

func Load() *Config {
    dir, err := os.Getwd()
    if err != nil {
        log.Fatalf("Error getting current working directory: %v", err)
    }

    envPath := filepath.Join(dir, "../.env")
    if err := godotenv.Load(envPath); err != nil {
        log.Fatalf("Error loading .env file from path %s: %v", envPath, err)
    }
    config := &Config{}

    config.HTTP_PORT = cast.ToString(coalesce("HTTP_PORT", ":8080"))
    config.DB_HOST = cast.ToString(coalesce("DB_HOST", "localhost"))
    config.DB_PORT = cast.ToInt(coalesce("DB_PORT", 5432))
    config.DB_USER = cast.ToString(coalesce("DB_USER", "postgres"))
    config.DB_PASSWORD = cast.ToString(coalesce("DB_PASSWORD", "dodi"))
    config.DB_NAME = cast.ToString(coalesce("DB_NAME", "www"))

    return config
}

func coalesce(key string, value interface{}) interface{} {
    val, exists := os.LookupEnv(key)
    if exists {
        return val
    }
    return value
}
