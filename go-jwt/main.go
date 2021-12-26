package main

import (
	"fmt"
	"os"

	"go_viliyadu/go-jwt/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("main func")

	port := os.Getenv("PORT")

	if port == "" {
		port = "8888"

		router := gin.New()
		router.Use(gin.Logger())

		routes.AuthRoutes(router)
		routes.UserRoutes(router)

		router.GET("/api-1", func(c *gin.Context) {
			c.JSON(200, gin.H{"success": "Access granted for api-1"})
		})

		router.GET("/api-2", func(c *gin.Context) {
			c.JSON(200, gin.H{"success": "Access granted for api-2"})
		})

		router.Run(":" + port)
	}
}
