package controllers

import (
	"context"
	"go_viliyadu/go-jwt/database"
	"go_viliyadu/go-jwt/models"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var validate = validator.New()

func HashPassword()

func VerifyPassword()

func Signup(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	count, err := userCollection.CountDocuments(ctx, bson.M{"email": user.Email})
	defer cancel()
	if err != nil {
		log.Panic(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while checking for the email"})
	}

	count, err = userCollection.CountDocuments(ctx, bson.M{"phone": user.Phone})
	defer cancel()
	if err != nil {
		log.Panic(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while checking for the Phone"})
	}

	if count > 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "this email or phone number already exists"})
	}
}

func Login()

func GetUsers()

func GetUser(c *gin.Context) {

}
