package routes

import (
	"go_viliyadu/go-ecom/controllers"

	"github.com/gin-gonic/gin"
)

func UserRouters(routes *gin.Engine) {

	routes.POST("/users/signup", controllers.SignUpCon())
	routes.POST("/users/login", controllers.LoginCon())
	routes.POST("/admin/addproduct", controllers.ProductViewerAdmin())
	routes.POST("/users/productview", controllers.SearchProduct())
	routes.POST("/users/search", controllers.SearchProductByQuery())

}
