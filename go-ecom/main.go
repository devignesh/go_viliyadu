package main

import (
	"fmt"
	"go_viliyadu/go-ecom/database"
	"go_viliyadu/go-ecom/middleware"
	"go_viliyadu/go-ecom/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("main func")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	app := controllers.NewApplication(database.ProductData(database.ClientMongo, "Products"), database.UserData(database.ClientMongo, "Users"))

	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRouters(router)
	router.Use(middleware.Authentications())

	router.GET("/addtocart", app.AddToCart())
	router.GET("/removeitem", app.RemoveItem())
	router.GET("/cartcheckout", app.BuyFromCart())
	router.GET("/instantbuy", app.InstantBuy())

	log.Fatal(router.Run(":" + port))
}
