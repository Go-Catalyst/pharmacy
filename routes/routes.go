package routes

import (
	user "pharmacy/internal/users/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, userHandler *user.UserHandler) {
	api := r.Group("/api")
	{
		api.GET("/users", userHandler.GetUsers)
		api.GET("/users/:id", userHandler.GetUser)
		api.POST("/users", userHandler.CreateUser)
		api.PUT("/users/:id", userHandler.UpdateUser)
		api.DELETE("/users/:id", userHandler.DeleteUser)
		api.POST("/jwt", userHandler.Login)

	}
}
