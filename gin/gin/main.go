package main

import (
	"gin/config"
	"gin/router"

	"github.com/gin-gonic/gin"
)

func main() {

	config.Connection()

	routers := gin.Default()

	router.Routes(routers)

	routers.Run("localhost:8080")
}
