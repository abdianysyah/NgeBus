package main

import (
	"github.com/gin-gonic/gin"
	"github.com/abdianysyah/backend/database"
)

func main()  {
	database.Connect()

	r := gin.Default()

	r.Run(":8080")
}