package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Configuration struct {
	Port       string
	DB         string
	DBHOST     string
	DBUSER     string
	DBPASSWORD string
	DBPORT     string
	DBNAME     string
}

var Config Configuration

func LoadConfig() {

	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	Config = Configuration{
		Port:       getEnv("PORT", "8080"),
		DB:         getEnv("DB", "pg"),
		DBHOST:     getEnv("DBHOST", ""),
		DBUSER:     getEnv("DBUSER", ""),
		DBPASSWORD: getEnv("DBPASSWORD", ""),
		DBPORT:     getEnv("DBPORT", ""),
		DBNAME:     getEnv("DBNAME", ""),
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
