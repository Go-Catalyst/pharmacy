package config

import (
    "log"
    "os"
)

type Configuration struct {
    Port string
}

var Config Configuration

func LoadConfig() {
    Config = Configuration{
        Port: getEnv("PORT", "8080"),
    }
}

func getEnv(key, defaultValue string) string {
    value, exists := os.LookupEnv(key)
    if !exists {
        log.Printf("Environment variable %s not set, using default value %s", key, defaultValue)
        return defaultValue
    }
    return value
}
