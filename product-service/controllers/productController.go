package controllers

import (
	"github.com/EgaTypeR/microservice-app/product-service/models"
	"github.com/EgaTypeR/microservice-app/product-service/utils"
	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	var products []models.Product
	utils.DB.Find(&products)
	c.JSON(200, gin.H{
		"products": products,
	})

}
func GetProduct(c *gin.Context) {

}
func GetBanner(c *gin.Context) {

}
func GetPromoProduct(c *gin.Context) {

}
