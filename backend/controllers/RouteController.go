package controllers


import (
	"net/http"

	"github.com/abdianysyah/backend/database"
	"github.com/abdianysyah/backend/models"
	"github.com/gin-gonic/gin"
)

func GetAllRoute(c *gin.Context)  {
	search := c.Query("search")

	var routes []models.Route

	query := database.DB.Model(&models.Route{})

	if search != "" {
		query = query.Where(
			"origin LIKE ? OR destination LIKE ?",
			"%"+search+"%",
			"%"+search+"%",
		)
		
	}

	query.Find(&routes)

	c.JSON(http.StatusOK, gin.H{
		"data" : routes,
	})
}


func CreateRoute(c *gin.Context)  {
	var routes models.Route

	if err := c.ShouldBindJSON(&routes); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : err.Error(),
		})
		return
	}

	database.DB.Create(&routes)

	c.JSON(http.StatusOK, gin.H{
		"message" : "Rute berhasil ditambahkan!",
		"data" : routes,
	})
}

func GetRouteByID(c *gin.Context)  {
	id := c.Param("id")

	var routes models.Route

	if err := database.DB.First(&routes, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error" : "Rute tidak ditemukan!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data" : routes,
	})
}


func UpdateRoute(c *gin.Context)  {
	id := c.Param("id")

	var routes models.Route

	if err := database.DB.First(&routes, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error" : "Rute tidak ditemukan!",
		})
		return
	}

	if err := c.ShouldBindJSON(&routes); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : err.Error(),
		})
	}

	database.DB.Save(&routes)

	c.JSON(http.StatusOK, gin.H{
		"message" : "Rute berhasil diupdate!",
		"rute" : routes,
	})
}

func DeleteRoute(c *gin.Context)  {
	id := c.Param("id")

	if err := database.DB.Delete(&models.Route{}, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : "Gagal menghapus rute",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message" : "Rute berhasil dihapus!",
	})
}