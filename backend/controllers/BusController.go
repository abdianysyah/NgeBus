package controllers

import (
	"net/http"

	"github.com/abdianysyah/backend/database"
	"github.com/abdianysyah/backend/models"
	"github.com/gin-gonic/gin"
)

// get All Bus
func GetAllBus(c *gin.Context)  {
	search := c.Query("search")

	var buses []models.Bus

	query := database.DB.Model(&models.Bus{})

	if search != "" {
		query = query.Where(
			"bus_name LIKE ? OR bus_number LIKE ?",
			"%"+search+"%",
			"%"+search+"%",
		)
	}
	query.Find(&buses)

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

	database.DB.Create(&bus)

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

