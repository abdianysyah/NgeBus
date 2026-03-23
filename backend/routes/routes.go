package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/abdianysyah/backend/controllers"
	"github.com/abdianysyah/backend/middlewares"
)

func UserRoutes(router *gin.Engine)  {
	auth := router.Group("/api")
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)
	}

	user := router.Group("/api/")
	user.Use(middlewares.AuthMiddleware())
	{
		user.GET("/user/dashboard", controllers.UserDashboard)
		user.GET("/bus", controllers.GetAllBus)
		user.GET("/bus/:id", controllers.GetBusByID)
		user.GET("/route", controllers.GetAllRoute)
		user.GET("/route/:id", controllers.GetRouteByID)
	}

	admin := router.Group("/api/")
	admin.Use(middlewares.AuthMiddleware(), middlewares.AdminOnly())
	{
		admin.GET("/admin/dashboard", controllers.AdminDashboard)
		admin.POST("/bus", controllers.CreateBus)
		admin.PUT("/bus/:id", controllers.UpdateBus)
		admin.DELETE("/bus/:id", controllers.DeleteBus)
		admin.POST("/route", controllers.CreateRoute)
		admin.PUT("/route/:id", controllers.UpdateRoute)
		admin.DELETE("/route/:id", controllers.DeleteRoute)
	}
}