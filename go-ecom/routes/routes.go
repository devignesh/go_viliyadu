package routes

import (
	"go-hotel/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(routes *gin.Engine) {

	routes.POST("/users/signup", controllers.SignUp())
	routes.POST("/users/login", controllers.Login())
	routes.POST("/admin/addproduct", controllers.ProductViewerAdmin())
	routes.POST("/users/productview", controllers.SearchProduct())
	routes.POST("/users/search", controllers.SearchProductByQuery())

}
