package controllers

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/EgaTypeR/microservice-app/product-service/models"
	"github.com/EgaTypeR/microservice-app/product-service/utils"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func GetProducts(c *gin.Context) {
	var products []models.Product
	err := utils.DB.Find(&products).Limit(10).Error
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"products": products,
	})

}
func GetProduct(c *gin.Context) {
	var product models.Product
	productID := c.Param("id")
	err := utils.DB.First(&product, productID).Error
	if err != nil {
		c.JSON(404, gin.H{
			"message": "Product not found",
		})
		return
	}
	c.JSON(200, gin.H{
		"product": product,
	})
}
func GetBanner(c *gin.Context) {

}
func GetPromoProduct(c *gin.Context) {

}

func GetActiveFlashSale(c *gin.Context) {
	var flashSaleProducts []models.FlashSaleProduct
	query := `
		SELECT * 
		FROM flash_sale
		JOIN product ON flash_sale.product_id = product.product_id
		WHERE start_date <= NOW() AND end_date >= NOW()
		LIMIT 10
	`

	ctx := context.Background()

	cachedFlashSales, err := utils.RedisClient.Get(ctx, "active_flash_sales").Result()
	if err == redis.Nil {
		err = utils.DB.Raw(query).Scan(&flashSaleProducts).Error
		if err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
			return
		}

		flashSaleProductsJSON, err := json.Marshal(flashSaleProducts)
		if err == nil {
			utils.RedisClient.Set(ctx, "active_flash_sales", flashSaleProductsJSON, time.Until(time.Now().Add(30*time.Minute)))
		}

		log.Println("Get data from database")
		c.JSON(200, gin.H{
			"data":    flashSaleProducts,
			"message": "Success get active flash sale products",
		})
	} else if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
	} else {
		json.Unmarshal([]byte(cachedFlashSales), &flashSaleProducts)
		log.Println("Get data from cache")
		c.JSON(200, gin.H{
			"data": flashSaleProducts,
		})
	}
}
