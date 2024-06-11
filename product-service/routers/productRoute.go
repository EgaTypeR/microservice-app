package routers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ProductRoute(db *gorm.DB, route *gin.Engine) {
	groupRoute := route.Group("/api/v1")
	groupRoute.GET("/get-products")
	groupRoute.GET("/show-product")
	groupRoute.GET("/baner-promo")
	groupRoute.GET("/promo-product")
	groupRoute.GET("/")

}
