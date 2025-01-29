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
    
    api.GET("/categories/:id", categoryHandlers.GetCategory)
    api.POST("/categories", categoryHandlers.CreateCategory)
    api.PUT("/categories/:id", categoryHandlers.UpdateCategory)
    api.DELETE("/categories/:id", categoryHandlers.DeleteCategory)

	}
  

}
