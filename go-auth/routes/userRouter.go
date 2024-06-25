package routes

import (
	controller "go-auth/controllers"
	"go-auth/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.GetUser())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())
}
