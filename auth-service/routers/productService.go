package routers

import (
	"github.com/EgaTypeR/microservice-app/auth-service/controllers"
	"github.com/EgaTypeR/microservice-app/auth-service/middlewares"
	"github.com/gin-gonic/gin"
)

func ProductRoute(r *gin.Engine) {
	groupRoute := r.Group("/api/v1", middlewares.AuthMiddleware())
	groupRoute.GET("/get-products", controllers.GetProducts)
	groupRoute.GET("/show-product/:id", controllers.GetProduct)
	groupRoute.GET("/baner-promo")
	groupRoute.GET("/promo-product")

	groupRoute.GET("/flash-sale/active", controllers.GetActiveFlashSale)
}
