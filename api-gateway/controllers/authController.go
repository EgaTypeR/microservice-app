package controllers

import (
	"context"
	"net/http"

	"github.com/EgaTypeR/microservice-app/auth-service/models"
	"github.com/EgaTypeR/microservice-app/auth-service/utils"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

	credentials := models.Credentials{}
	user := models.RegisterUser{}

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := utils.MySqlDB.Where("email = ? ", credentials.Email).First(&user).Error
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	if !utils.ComparePassword(user.Password, []byte(credentials.Password)) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := utils.GenerateToken(user.FirstName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		return
	}

	user.Password = ""

	c.JSON(http.StatusOK, gin.H{
		"token":   token,
		"user":    user,
		"message": "Login successful",
	})

	// Login Controller
}
func Register(c *gin.Context) {
	// Register Controller
	var user models.RegisterUser
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if user.FirstName == "" || user.LastName == "" || user.Email == "" || user.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "All fields are required"})
		return
	}

	user.Password = utils.HashPassword([]byte(user.Password))

	err := utils.MySqlDB.Create(&user).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User created successfully go to login page to login",
	})
}

func Logout(c *gin.Context) {
	// Logout Controller
	tokenString := c.Query("token")
	if tokenString == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Missing token"})
		return
	}
	claims, err := utils.ValidateToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
		return
	}
	jti := claims.ID
	ctx := context.Background()
	err = utils.RedisClient.Del(ctx, jti).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error logging out"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})

}
