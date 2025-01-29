package main

import (
	"fmt"
	"log"
	"pharmacy/config"
	"pharmacy/docs"
	userHandlers "pharmacy/internal/users/handlers"
	userRepo "pharmacy/internal/users/repository"
	"pharmacy/routes"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title User CRUD API
// @version 1.0
// @description This is a sample server for a user CRUD API.
// @host localhost:8080
// @BasePath /

func main() {
	cfg := config.LoadConfig()

	if cfg.DBType == "PG" {
		config.InitPGDB(cfg.DBPath)
	} else {
		// config.InitDB(cfg.DBPath)
		fmt.Println("not work with sqlite!")
	}

	// Set Gin mode to "release" to disable debug logs
	gin.SetMode(gin.ReleaseMode)

	// Create Gin router
	r := gin.Default()


	// Set up repository and handler
	userRepo := userRepo.NewUserRepository(config.DB)
	userHandler := userHandlers.NewUserHandler(userRepo)

	// Set up routes
	routes.SetupRoutes(r, userHandler)

	// Initialize Swagger documentation
	docs.SwaggerInfo.BasePath = "/api"

	// Set up Swagger UI
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Start the server
	log.Println("Server is running on port", cfg.Port)

	// Log the Swagger URL before running the server
	log.Printf("Swagger url: http://localhost:%s/swagger/index.html", cfg.Port)

	// Run the Gin server
	log.Println(r.Run(":" + cfg.Port))
}
