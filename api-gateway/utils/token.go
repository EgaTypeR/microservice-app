package utils

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/EgaTypeR/microservice-app/auth-service/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func GenerateToken(username string) (string, error) {
	jwtKey := os.Getenv("JWT_SECRET")
	if jwtKey == "" {
		return "", errors.New("JWT_SECRET is not set")
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &models.Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ID:        uuid.New().String(), // Generate a unique ID
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		return "", err
	}

	// Store the token ID (jti) in Redis
	ctx := context.Background()
	err = RedisClient.Set(ctx, claims.ID, "active", time.Until(expirationTime)).Err()
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(tokenString string) (*models.Claims, error) {
	jwtKey := os.Getenv("JWT_SECRET")
	if jwtKey == "" {
		return nil, errors.New("JWT_SECRET is not set")
	}

	claims := &models.Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("Invalid token")
	}

	// Check if the token is blacklisted
	ctx := context.Background()
	_, err = RedisClient.Get(ctx, claims.ID).Result()
	if err != nil {
		return nil, errors.New("Token is blacklisted")
	}

	return claims, nil
}
