package main

import (
	"gin/config"
	"gin/router"

	"github.com/gin-gonic/gin"
)

func main() {

	fmt.Println("main func")


	config.Connection()

	routers := gin.Default()

	router.Routes(routers)

	routers.Run("localhost:8080")
}
