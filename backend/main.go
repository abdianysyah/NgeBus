package main

import (
	"time"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/abdianysyah/backend/database"
	"github.com/abdianysyah/backend/routes"

)

func main()  {
	database.Connect()

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge: 12 * time.Hour,
	}))
	
	routes.UserRoutes(router)

	router.Run(":8080")
}