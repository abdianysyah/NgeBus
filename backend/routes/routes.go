package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/abdianysyah/backend/controllers"
	"github.com/abdianysyah/backend/middlewares"
)

func UserRoutes(router *gin.Engine)  {
	api := router.Group("/api")
	{
		api.POST("/register", controllers.Register)
		api.POST("/login", controllers.Login)
	}

	user := router.Group("/api/user")
	user.Use(middlewares.AuthMiddleware())
	{
		user.GET("/dashboard", controllers.UserDashboard)
	}

	admin := router.Group("/api/admin")
	admin.Use(middlewares.AuthMiddleware(), middlewares.AdminOnly())
	{
		admin.GET("/dashboard", controllers.AdminDashboard)
	}
}