package middlewares

import (
	"context"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/EgaTypeR/microservice-app/auth-service/models"
	"github.com/EgaTypeR/microservice-app/auth-service/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
)

func AuthMiddleware() gin.HandlerFunc {
	var jwtKey = []byte(os.Getenv("JWT_SECRET"))
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Missing token"})
			c.Abort()
			return
		}
		splitToken := strings.Split(authHeader, "Bearer ")
		if len(splitToken) != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token format"})
			c.Abort()
			return
		}

		tokenString := splitToken[1]
		tokenString = strings.TrimSpace(tokenString) // Trim any leading/trailing whitespace
		log.Println("Token:", tokenString)

		claims := &models.Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		log.Println(token)

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token signature"})
				c.Abort()
				return
			}
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
			c.Abort()
			return
		}

		if !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
			c.Abort()
			return
		}

		jti := claims.ID
		ctx := context.Background()
		isBlacklisted, err := utils.RedisClient.Get(ctx, jti).Result()
		if err != nil && err != redis.Nil {
			log.Printf("Error checking token in Redis: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error checking token"})
			c.Abort()
			return
		} else if isBlacklisted == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Token is not active"})
			c.Abort()
			return
		}

		c.Set("username", claims.Username)
		c.Next()
	}
}
