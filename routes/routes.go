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

	r.POST("/categories", categoryHandlers.CreateCategory)
}
