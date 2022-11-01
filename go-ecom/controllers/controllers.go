package controllers

import "github.com/gin-gonic/gin"

func HashPassword(password string) string {}

func VerifyPassword(userPassword, givenPassword string) (bool, string) {}

func SignUpCon() gin.HandlerFunc {

}

func LoginCon() gin.HandlerFunc {}

func ProductViewerAdmin() gin.HandlerFunc {}

func SearchProduct() gin.HandlerFunc {}

func SearchProductByQuery() gin.HandlerFunc {}
