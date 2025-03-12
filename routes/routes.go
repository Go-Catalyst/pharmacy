package routes

import (
	category "pharmacy/internal/categories/handlers"
	drug "pharmacy/internal/drugs/handlers"
	user "pharmacy/internal/users/handlers"
	jwt "pharmacy/internal/users/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, userHandler *user.UserHandler, categoryHandlers *category.CategoryHandler, drugHandler *drug.DrugHandler) {
	api := r.Group("/api")
	
		user := api.Group("/users")
		
		user.GET("", userHandler.GetUsers)
		user.GET("/:id", userHandler.GetUser, jwt.AuthMiddleware())
		user.POST("", userHandler.CreateUser)
		user.PUT("/:id", userHandler.UpdateUser, jwt.AuthMiddleware())
		user.DELETE("/:id", userHandler.DeleteUser, jwt.AuthMiddleware())
		user.POST("/jwt", userHandler.Login)

		category := api.Group("/categories", jwt.AuthMiddleware() ) 
		
		category.GET("/:id", categoryHandlers.GetCategory)
		category.POST("", categoryHandlers.CreateCategory)
		category.PUT("/:id", categoryHandlers.UpdateCategory)
		category.DELETE("/:id", categoryHandlers.DeleteCategory)
	
		drug := api.Group("/drugs", jwt.AuthMiddleware())

		drug.GET("", drugHandler.GetAllDrugs)    
		drug.GET("/:id", drugHandler.GetDrugByID)  
		drug.POST("", drugHandler.AddDrug)

	

}
