package routers

import (
	"github.com/EgaTypeR/microservice-app/auth-service/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoute(route *gin.Engine) {
	groupRoute := route.Group("/api/v1")
	groupRoute.POST("/register", controllers.Register)
	groupRoute.POST("/login", controllers.Login)
	groupRoute.POST("/logout", controllers.Logout)
	groupRoute.POST("/refresh")
	groupRoute.POST("/validate")
}
