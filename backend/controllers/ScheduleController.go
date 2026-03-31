package controllers

import (
	"net/http"

	"github.com/abdianysyah/backend/database"
	"github.com/abdianysyah/backend/models"
	"github.com/gin-gonic/gin"
)

func GetAllSchedule(c *gin.Context) {
	var schedules []models.Schedule

	if err := database.DB.
		Preload("Bus").
		Preload("Bus.Company").
		Preload("Route").
		Find(&schedules).Error; err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Gagal memuat data!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": schedules,
	})
}

func CreateSchedule(c *gin.Context)  {
	var schedule models.Schedule

	if err := c.ShouldBindJSON(&schedule); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : err.Error(),
		})
		return
	}

	database.DB.Create(&schedule)

	c.JSON(http.StatusOK, gin.H{
		"message" : "Jadwal berhasil ditambahkan!",
		"data" : schedule,
	})
}

func UpdateSchedule(c *gin.Context)  {
	id := c.Param("id")

	var schedule models.Schedule

	if err := database.DB.First(&schedule, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error" : "Jadwal tidak ditemukan!",
		})
		return
	}

	if err := c.ShouldBindJSON(&schedule); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : err.Error(),
		})
		return
	}

	database.DB.Save(&schedule)

	c.JSON(http.StatusOK, gin.H{
		"message" : "Jadwal berhasil diperbarui!",
		"data" : schedule,
	})
}

func DeleteSchedule(c *gin.Context)  {
	id := c.Param("id")

	if err := database.DB.Delete(&models.Schedule{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error" : "Jadwal tidak ditemukan!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message" : "Jadwal berhasil dihapus!",
	})
}