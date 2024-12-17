package main

import (
    "github.com/gin-gonic/gin"
    _ "pharmacy/docs"
    "github.com/swaggo/gin-swagger"
    "github.com/swaggo/files"
    "pharmacy/config"
    "pharmacy/routes"
)

// @title User CRUD API
// @version 1.0
// @description This is a sample server for a user CRUD API.
// @host localhost:8080
// @BasePath /

func main() {
    config.LoadConfig()

    r := gin.Default()
    routes.RegisterRoutes(r)

    // Swagger UI
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    r.Run(":" + config.Config.Port)
}
