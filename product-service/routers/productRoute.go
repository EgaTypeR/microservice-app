package routers

import (
	"github.com/EgaTypeR/microservice-app/product-service/controllers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ProductRoute(db *gorm.DB, route *gin.Engine) {
	groupRoute := route.Group("/api/v1")
	groupRoute.GET("/get-products", controllers.GetProducts)
	groupRoute.GET("/show-product/:id", controllers.GetProduct)
	groupRoute.GET("/baner-promo")
	groupRoute.GET("/promo-product")

	groupRoute.GET("/flash-sale/active", controllers.GetActiveFlashSale)
}
