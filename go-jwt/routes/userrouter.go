package routes

import (
	"go_viliyadu/go-jwt/controllers"
	"go_viliyadu/go-jwt/middleware"

	"github.com/gin-gonic/gin"
)

//user routes func
func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate)
	incomingRoutes.GET("/users", controllers.GetUsers)
	incomingRoutes.GET("/users/:user_id", controllers.GetUser)
}
