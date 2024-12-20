package routes

import (
	"github.com/gin-gonic/gin"
	handlers2 "pharmacy/handlers"
	"pharmacy/middelware"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/users", handlers2.GetUsers)
	r.GET("/users/:id", handlers2.GetUser)
	r.POST("/users", handlers2.CreateUser)
	r.PUT("/users/:id", handlers2.UpdateUser)
	r.DELETE("/users/:id", handlers2.DeleteUser)

	r.POST("/categories", handlers2.CreateCategory)

	protected := r.Group("/protected")
	protected.Use(middelware.LoginAuth())
	protected.GET("/getUsers", handlers2.GetUsers)

	r.POST("/login", handlers2.Login)
}
