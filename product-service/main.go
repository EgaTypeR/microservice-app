package main

import (
	"log"
	"os"

	"github.com/EgaTypeR/microservice-app/product-service/routers"
	"github.com/EgaTypeR/microservice-app/product-service/utils"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error Loading .env file")
	}

	utils.InitiateDB()
	utils.SetupRedis()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	routers.ProductRoute(utils.DB, r)
	log.Fatal(r.Run(":" + os.Getenv("PORT")))
}
