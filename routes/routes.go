package routes

import (
	"github.com/gin-gonic/gin"
	categoryHandlers "pharmacy/internal/categories/handlers"
	userHandlers "pharmacy/internal/users/handlers"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/users", userHandlers.GetUsers)
	r.GET("/users/:id", userHandlers.GetUser)
	r.POST("/users", userHandlers.CreateUser)
	r.PUT("/users/:id", userHandlers.UpdateUser)
	r.DELETE("/users/:id", userHandlers.DeleteUser)
	r.GET("/categories", categoryHandlers.GetCategories)
	r.GET("/categories/:id", categoryHandlers.GetCategory)
	r.POST("/categories", categoryHandlers.CreateCategory)
	r.PUT("/categories/:id", categoryHandlers.UpdateCategory)
	r.DELETE("/categories/:id", categoryHandlers.DeleteCategory)
}
