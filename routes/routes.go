package routes

import (
	"github.com/gin-gonic/gin"
	category "pharmacy/internal/categories/handlers"
	drug "pharmacy/internal/drugs/handlers"
	user "pharmacy/internal/users/handlers"
)

func SetupRoutes(r *gin.Engine, userHandler *user.UserHandler, categoryHandlers *category.CategoryHandler, drugHandler *drug.DrugHandler) {
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

		api.POST("/drugs", drugHandler.AddDrug)
	}

}
