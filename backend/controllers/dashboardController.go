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
	var totalCompany int64

	database.DB.Model(&models.User{}).Count(&totalUsers)
	database.DB.Model(&models.Order{}).Count(&totalOrders)
	database.DB.Model(&models.Bus{}).Count(&totalBus)
	database.DB.Model(&models.Schedule{}).Count(&totalSchedules)
	database.DB.Model(&models.Company{}).Count(&totalCompany)

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

	type StatusOrder struct {
		Pending int64 `json:"pending"`
		Berhasil int64 `json:"berhasil"`
		Gagal int64 `json:"gagal"`
	}

	var statusOrder StatusOrder

	database.DB.Model(&models.Order{}).
		Where("status = ?", "Pending").
		Count(&statusOrder.Pending)

	database.DB.Model(&models.Order{}).
		Where("status = ?", "Berhasil").
		Count(&statusOrder.Berhasil)

	database.DB.Model(&models.Order{}).
		Where("status = ?", "Gagal").
		Count(&statusOrder.Gagal)

	c.JSON(http.StatusOK, gin.H{
		"message": "Admin dashboard data",
		"data": gin.H{
			"total_users": totalUsers,
			"total_orders": totalOrders,
			"total_bus": totalBus,
			"total_schedules": totalSchedules,
			"total_company": totalCompany,
			"monthly_orders": monthlyOrders,
			"status_order": statusOrder,
		},
	})
}
