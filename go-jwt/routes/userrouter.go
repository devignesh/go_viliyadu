package routes

import (
	"go_viliyadu/go-jwt/controllers"

	"github.com/gin-gonic/gin"
	"github.com/go-chi/chi/middleware"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controllers.GetUsers())
	incomingRoutes.GET("/users/:user_id", controllers.GetUser())
}
