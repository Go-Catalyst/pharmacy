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
		DB:         getEnv("DB", "sqllite"),
		DBHOST:     getEnv("DBHOST", "localhost"),
		DBUSER:     getEnv("DBUSER", "postgres"),
		DBPASSWORD: getEnv("DBPASSWORD", "password"),
		DBPORT:     getEnv("DBPORT", "5432"),
		DBNAME:     getEnv("DBNAME", "phdb"),
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
