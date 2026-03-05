package controllers

import (
	"net/http"

	"github.com/abdianysyah/backend/database"
	"github.com/abdianysyah/backend/models"
	"github.com/abdianysyah/backend/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context)  {
	var input struct {
		Name		string `json:"name"`
		Email		string `json:"email"`
		Phone		string `json:"phone"`
		Password	string `json:"password"`
	}	

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), 14)

	user := models.User{
		Name:	input.Name,
		Email:	input.Email,
		Phone:	input.Phone,
		Password: string(hashPassword),
		Role:	"user",
	}

	database.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{
		"message" : "Register Berhasil",
	})
}

func Login(c *gin.Context)  {
	var input struct {
		Email		string	`json:"email"`
		Password	string	`json:"password"`
	}

	var user models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Where("email = ?", input.Email).First(&user)

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message" : "Password salah",
		})		
	}

	token, err := utils.GenerateToken(user.ID.String(), user.Role)

	if err != nil {
		c.JSON(500, gin.H{
			"error" : "failed generate token",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message" : "Login berhasil",
		"role" : user.Role,
		"token" : token,
	})
}