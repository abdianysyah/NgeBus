package controllers

import (
	"net/http"

	"github.com/abdianysyah/backend/database"
	"github.com/abdianysyah/backend/models"
	"github.com/gin-gonic/gin"
)

func GetAllTicket(c *gin.Context)  {
	var tickets []models.Ticket

	c.JSON(http.StatusOK, gin.H{
		"data" : tickets,
	})
}

func CreateTicket(c *gin.Context)  {
	var ticket models.Ticket

	if err := c.ShouldBindJSON(&ticket); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : err.Error(),
		})
		return
	}

	database.DB.Create(&ticket)

	c.JSON(http.StatusOK, gin.H{
		"message" : "Tiket berhasil ditambahkan!",
		"data" : ticket,
	})
}

func UpdateTicket(c *gin.Context)  {
	id := c.Param("id")

	var ticket models.Ticket

	if err := database.DB.First(&ticket, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error" : "Tiket tidak ditemukan!",
		})
		return
	}

	if err := c.ShouldBindJSON(&ticket); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : err.Error(),
		})
		return
	}

	database.DB.Save(&ticket)

	c.JSON(http.StatusOK, gin.H{
		"message" : "Tiket berhasil diperbarui!",
		"data" : ticket,
	})
}

func DeleteTicket(c *gin.Context)  {
	id := c.Param("id")

	if err := database.DB.Delete(&models.Ticket{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error" : "Tiket tidak ditemukan!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message" : "Tiket berhasil dihapus!",
	})
}