package main

import (
	"log"
	"os"

	"github.com/EgaTypeR/microservice-app/auth-service/middlewares"
	"github.com/EgaTypeR/microservice-app/auth-service/routers"
	"github.com/EgaTypeR/microservice-app/auth-service/utils"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error Loading .env file")
	}
	r := gin.Default()
	utils.SetupRedis()
	utils.SetupMysql()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	routers.AuthRoute(r)

	productService := os.Getenv("PRODUCT_SERVICE_URL")
	log.Println("Product Service URL:", productService)
	authenticated := r.Group("/")
	authenticated.Use(middlewares.AuthMiddleware())
	{
		authenticated.Any("/product/*proxyPath", utils.ReverseProxy(os.Getenv("PRODUCT_SERVICE_URL")))
	}

	log.Fatal(r.Run(":" + os.Getenv("PORT")))
}
