package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/abdianysyah/backend/controllers"
	"github.com/abdianysyah/backend/middlewares"
)

func SetupRoutes(router *gin.Engine)  {
	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)

	user := router.Group("/user")
	user.Use(middlewares.AuthMiddleware())
	{
		// user.GET("/profile", controllers.Por)
	}

	admin := router.Group("/admin")
	admin.Use(middlewares.AuthMiddleware(), middlewares.AdminOnly())
	{
		// admin.GET("/dashboard", )
	}
}