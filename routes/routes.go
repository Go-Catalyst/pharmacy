package routes

import (
	"github.com/gin-gonic/gin"
	"pharmacy/internal/categories/category_handlers"
	"pharmacy/internal/users/handlers"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/users", handlers.GetUsers)
	r.GET("/users/:id", handlers.GetUser)
	r.POST("/users", handlers.CreateUser)
	r.PUT("/users/:id", handlers.UpdateUser)
	r.DELETE("/users/:id", handlers.DeleteUser)

	r.POST("/categories", category_handlers.CreateCategory)
}
