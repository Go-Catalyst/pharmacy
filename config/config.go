package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Configuration struct {
	DBPath    string
	DBType    string
	Username  string
	Password  string
	APIDomain string
	URIMongo  string
	GoSecret  string
	URIRabbit string
	Port      string
}

var Config Configuration

func LoadConfig() *Configuration {
	err := godotenv.Load()
	if err != nil {

		log.Fatalf("Error loading .env file: %v", err)
	}

	return &Configuration{
		DBPath:    os.Getenv("DB_PATH"),
		DBType:    os.Getenv("DB_TYPE"),
		Username:  os.Getenv("USST"),
		Password:  os.Getenv("PSST"),
		APIDomain: os.Getenv("API_DOMAIN"),
		URIMongo:  os.Getenv("URI_MONGO"),
		GoSecret:  os.Getenv("GOLANG_SECRET"),
		URIRabbit: os.Getenv("URI_RABBIT"),
		Port:      os.Getenv("PORT"),
	}
}
