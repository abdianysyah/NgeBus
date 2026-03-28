package controllers

import (
	"net/http"

	"github.com/abdianysyah/backend/database"
	"github.com/abdianysyah/backend/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// get All Bus
func GetAllBus(c *gin.Context)  {
	search := c.Query("search")
	status := c.Query("status")
	classes := c.Query("bus_class")

	var buses []models.Bus

	query := database.DB.Model(&models.Bus{})

	if search != "" {
		query = query.Where(
			"bus_name LIKE ? OR bus_number LIKE ?",
			"%"+search+"%",
			"%"+search+"%",
		)
	}

	if status != "" {
		query = query.Where("status = ?", status)
	}

	if classes != "" {
		query = query.Where("bus_class = ?", classes)
	}

	query.Preload("Company").Find(&buses)

	c.JSON(http.StatusOK, gin.H{
		"data" : buses,
	})
}

// Tambah Data Bus
func CreateBus(c *gin.Context)  {
	var bus models.Bus

	if err := c.ShouldBindJSON(&bus); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : err.Error(),
		})
		return
	}

	if err := database.DB.Create(&bus).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error" : "Gagal menambahkan bus!",
		})
		return
	}

	if err := database.DB.Model(&models.Company{}).
		Where("id = ?", bus.CompanyID).
		Update("total_bus", gorm.Expr("total_bus + ?", 1)).Error; err != nil {
		
		c.JSON(http.StatusInternalServerError, gin.H{
			"error" : "Bus berhasil dibuat, tapi gagal diupdate total bus!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message" : "Bus berhasil ditambahkan!",
		"data" : bus,
	})
}

// Get Data Bus by ID
func GetBusByID(c *gin.Context)  {
	id := c.Param("id")

	var bus models.Bus

	if err := database.DB.First(&bus, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error" : "Bus tidak ditemukan",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data" : bus,
	})
}

// Update Data Bus
func UpdateBus(c *gin.Context)  {
	id := c.Param("id")

	var bus models.Bus

	if err := database.DB.First(&bus, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error" : "Bus tidak ditemukan",
		})
		return
	}

	if err := c.ShouldBindJSON(&bus); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : err.Error(),
		})
		return
	}

	database.DB.Save(&bus)

	c.JSON(http.StatusOK, gin.H{
		"message" : "Bus berhasil di update!",
		"data" : bus,
	})
}

func DeleteBus(c *gin.Context)  {
	id := c.Param("id")

	if err := database.DB.Delete(&models.Bus{}, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : "Gagal menghapus bus",
		})
		return

	}

	c.JSON(http.StatusOK, gin.H{
		"message" : "Bus berhasil dihapus!",
	})
}

