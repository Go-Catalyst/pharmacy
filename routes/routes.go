package routes

import (
    "github.com/gin-gonic/gin"
    "pharmacy/internal/users/handlers"
    "pharmacy/internal/categories/handlers"

)

func RegisterRoutes(r *gin.Engine) {
    r.GET("/users", handlers.GetUsers)
    r.GET("/users/:id", handlers.GetUser)
    r.POST("/users", handlers.CreateUser)
    r.PUT("/users/:id", handlers.UpdateUser)
    r.DELETE("/users/:id", handlers.DeleteUser)
    r.GET("/categories", handlers.GetCategories)
    r.GET("/categories/:id", handlers.GetCategory)
    r.POST("/categories", handlers.CreateCategory)
    r.PUT("/categories/:id", handlers.UpdateCategory)
    r.DELETE("/categories/:id", handlers.DeleteCategory)
}
