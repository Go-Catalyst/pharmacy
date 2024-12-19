package main

import (
	"pharmacy/config"
	"pharmacy/db"
	_ "pharmacy/docs"
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
	config.LoadConfig()
	db.Init()

	r := gin.Default()
	routes.RegisterRoutes(r)

	// Swagger UI
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err := r.Run(":" + config.Config.Port)
	if err != nil {
		panic("Error starting server: " + err.Error())
	}
}
