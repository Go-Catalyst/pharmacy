package config

import (
	categoryModels "pharmacy/internal/categories/models"
	"pharmacy/internal/users/models"

	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

// func InitDB(dbPath string) {
// 	// Check if the directory for the database file exists
// 	dir := getDirectory(dbPath)
// 	if _, err := os.Stat(dir); os.IsNotExist(err) {
// 		if err := os.MkdirAll(dir, 0755); err != nil {
// 			log.Fatalf("Failed to create directory for database: %v", err)
// 		}
// 	}

// 	// Check if the database file exists; create if it doesn't
// 	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
// 		file, err := os.Create(dbPath)
// 		if err != nil {
// 			log.Fatalf("Failed to create database file: %v", err)
// 		}
// 		file.Close() // Close the file after creation
// 		log.Println("Database file created")
// 	}

// 	// Initialize GORM connection
// 	var err error
// 	DB, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
// 	if err != nil {
// 		log.Fatalf("Failed to connect to database: %v", err)
// 	}
// 	log.Println("Database connected successfully")

// 	RunMigrations()
// }

// Helper function to get the directory of the file path
func getDirectory(path string) string {
	if idx := len(path) - 1; idx >= 0 && os.IsPathSeparator(path[idx]) {
		return path
	}
	for i := len(path) - 1; i >= 0; i-- {
		if os.IsPathSeparator(path[i]) {
			return path[:i]
		}
	}
	return "."
}

func InitPGDB(dsn string) {
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	log.Println("PostgreSQL database connected successfully")

	RunMigrations()
}

func RunMigrations() {
	// Run migrations
	err = DB.AutoMigrate(&models.User{})
	err = DB.AutoMigrate(&categoryModels.Category{})

	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}
	log.Println("Database migrated successfully")
}
