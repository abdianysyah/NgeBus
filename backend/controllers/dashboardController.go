package controllers

import (
	"net/http"
	"time"

	"github.com/abdianysyah/backend/database"
	"github.com/abdianysyah/backend/models"
	"github.com/gin-gonic/gin"
)

func UserDashboard(c *gin.Context)  {
	userID := c.MustGet("user_id")

	var totalOrders int64

	database.DB.Model(&models.Order{}).
		Where("user_id = ?", userID).
		Count(&totalOrders)
	
	c.JSON(http.StatusOK, gin.H{
		"message" : "User Dashboard data",
		"data" : gin.H{
			"total_orders" : totalOrders,
		},
	})
}

func AdminDashboard(c *gin.Context)  {
	var totalUsers int64
	var totalOrders int64
	var totalBus int64
	var totalSchedules int64

	database.DB.Model(&models.User{}).Count(&totalUsers)
	database.DB.Model(&models.Order{}).Count(&totalOrders)
	database.DB.Model(&models.Bus{}).Count(&totalBus)
	database.DB.Model(&models.Schedule{}).Count(&totalSchedules)

	currentYear := time.Now().Year()

	type Result struct {
		Month int
		Total int
	}

	var result []Result

	database.DB.Model(&models.Order{}).
		Select("MONTH(created_at) as month, COUNT(*) as total").
		Where("YEAR(created_at) = ?", currentYear).
		Group("MONTH(created_at)").
		Scan(&result)

	monthlyOrders := make([]int, 12)

	for _, r := range result {
		monthlyOrders[r.Month-1] = r.Total
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Admin dashboard data",
		"data": gin.H{
			"total_users": totalUsers,
			"total_orders": totalOrders,
			"total_bus": totalBus,
			"total_schedules": totalSchedules,
			"monthly_orders": monthlyOrders,
		},
	})
}
