package controllers

import (
	"net/http"

	"github.com/abdianysyah/backend/database"
	"github.com/abdianysyah/backend/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func GetUserByID()  {
	
}

func UpdateProfile(c *gin.Context)  {
	id := c.Param("id")

	var user models.User

	if err := database.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error" : "User tidak ditemukan",
		})
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : err.Error(),
		})
		return
	}

	database.DB.Save(&user)

	c.JSON(http.StatusOK,  gin.H{
		"message" : "User Berhasil di update!",
		"data" : user,
	})
}

func UpdatePassword(c *gin.Context)  {
	id := c.Param("id")

	var user models.User

	var input struct {
		Password	string `json:"password"`
	}

	if err := database.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error" : "User tidak ditemukan",
		})
		return
	}

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), 14)

	password := models.User {
		Password: string(hashPassword),
	}

	database.DB.Create(&password)
}

func DeleteUser(c *gin.Context)  {
	id := c.Param("id")

	if err := database.DB.Delete(&models.User{}, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : "Gagal Menghapus user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message" : "User berhasil dihapus!",
	})
}