package routes

import (
    "github.com/gin-gonic/gin"
    "pharmacy/internal/users"
    "pharmacy/internal/categories"

)

func RegisterRoutes(r *gin.Engine) {
    r.GET("/users", users.handlers.GetUsers)
    r.GET("/users/:id", users.handlers.GetUser)
    r.POST("/users", users.handlers.CreateUser)
    r.PUT("/users/:id", users.handlers.UpdateUser)
    r.DELETE("/users/:id", users.handlers.DeleteUser)
    r.GET("/categories", categories.handlers.GetCategories)
    r.GET("/categories/:id", categories.handlers.GetCategory)
    r.POST("/categories", categories.handlers.CreateCategory)
    r.PUT("/categories/:id", categories.handlers.UpdateCategory)
    r.DELETE("/categories/:id", categories.handlers.DeleteCategory)
}
