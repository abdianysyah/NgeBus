package controllers

import (
	"net/http"

	"github.com/abdianysyah/backend/database"
	"github.com/abdianysyah/backend/models"
	"github.com/gin-gonic/gin"
) 

func GetAllCompany(c *gin.Context)  {
	search := c.Query("search")

	var company []models.Company

	query := database.DB.Model(&models.Company{})

	if search != "" {
		query = query.Where(
			"name_company LIKE ? ","%"+search+"%",
		)
	}

	query.Find(&company)

	c.JSON(http.StatusOK, gin.H{
		"data" : company,
	})
}

func CreateCompany(c *gin.Context)  {
	var company models.Company

	if err := c.ShouldBindJSON(&company); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : err.Error(),
		})
		return		
	}

	database.DB.Create(&company)

	c.JSON(http.StatusOK, gin.H{
		"message" : "PO Bus berhasil ditambahkan!",
		"data" : company,
	})
}

func GetCompanyByID(c *gin.Context)  {
	id := c.Param("id")

	var company models.Company

	if err := database.DB.First(&company, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error" : "PO Bus tidak ditemukan!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data" : company,
	})
}

func UpdateCompany(c *gin.Context)  {
	id := c.Param("id")

	var company models.Company

	if err := database.DB.First(&company, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error" : "PO Bus tidak ditemukan!",
		})
		return
	}

	var input models.Company
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : err.Error(),
		})
		return
	}

	company.NameCompany = input.NameCompany
	company.TotalBus = input.TotalBus

	database.DB.Save(&company)

	c.JSON(http.StatusOK, gin.H{
		"message" : "PO Bus berhasil diupdate!",
		"data" : company,
	})
}

func DeleteCompany(c *gin.Context)  {
	id := c.Param("id")

	if err := database.DB.Delete(&models.Company{}, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : "Gagal Menghapus PO Bus",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message" : "PO Bus berhasil dihapus!",
	})
}