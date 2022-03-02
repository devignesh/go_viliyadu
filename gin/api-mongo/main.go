package main

import (
	"fmt"
	"go_viliyadu/gin/api-mongo/handler"

	"github.com/gin-gonic/gin"
)

//main func
func main() {
	fmt.Println("main")

	router := gin.Default()

	router.POST("/food", handler.FoodCreate)
	router.Run("localhost:8080")
}
