package controllers

import (
	"net/http"

	"github.com/EgaTypeR/microservice-app/auth-service/models"
	"github.com/EgaTypeR/microservice-app/auth-service/utils"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

	credentials := models.Credentials{}
	user := models.User{}

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := utils.MySqlDB.Where("email = ? AND password = ?", credentials.Email, credentials.Password).First(&user).Error
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := utils.GenerateToken(user.FirstName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token":   token,
		"user":    user,
		"message": "Login successful",
	})

	// Login Controller
}
func Register(c *gin.Context) {
	// Register Controller
}

func Logout(c *gin.Context) {
	// Logout Controller
}
