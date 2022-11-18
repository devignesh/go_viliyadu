package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"go_viliyadu/go-ecom/database"
	dto "go_viliyadu/go-ecom/models"
	tokengen "go_viliyadu/go-ecom/token"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var UserCollection *mongo.Collection = database.UserData(database.ClientMongo, "Users")
var ProductCollection *mongo.Collection = database.ProductData(database.ClientMongo, "Products")
var validate = validator.New()

func HashPassword(password string) string {

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}

func VerifyPassword(userPassword, givenPassword string) (bool, string) {

	err := bcrypt.CompareHashAndPassword([]byte(givenPassword), []byte(userPassword))
	valid := true
	msg := ""
	if err != nil {
		msg = "Login Or Passowrd is Incorerct"
		valid = false
	}
	return valid, msg
}

func SignUpCon() gin.HandlerFunc {

	return func(c *gin.Context) {

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var user dto.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := validate.Struct(user)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr})
			return
		}

		count, err := UserCollection.CountDocuments(ctx, bson.M{"email": user.Email})
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User already exists"})
		}

		count, err = UserCollection.CountDocuments(ctx, bson.M{"phone": user.Phone})
		defer cancel()

		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Phone is already in use"})
			return
		}

		password := HashPassword(*user.Password)
		user.Password = &password

		user.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.ID = primitive.NewObjectID()
		user.UserID = user.ID.Hex()

		token, refreshtoken, _ := tokengen.TokenGenerator(*user.Email, *user.FistName, *user.LastName, user.UserID)

		user.Token = &token
		user.RefreshToken = &refreshtoken
		user.UserCart = make([]dto.ProductUser, 0)
		user.AddressDetails = make([]dto.Address, 0)
		user.OrderStatus = make([]dto.Order, 0)

		_, inserterr := UserCollection.InsertOne(ctx, user)
		if inserterr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "not created"})
			return
		}
		defer cancel()

		c.JSON(http.StatusCreated, "Successfully Signed Up!!")

	}

}

func LoginCon() gin.HandlerFunc {

	return func(c *gin.Context) {

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var loginuser dto.User
		var foundUser dto.User

		if err := c.BindJSON(&loginuser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		err := UserCollection.FindOne(ctx, bson.M{"email": loginuser.Email}).Decode(&foundUser)
		defer cancel()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "login or password incorrect"})
			return
		}

		PasswordCheck, msg := VerifyPassword(*loginuser.Password, *foundUser.Password)
		defer cancel()

		if !PasswordCheck {
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			fmt.Println(msg)
			return
		}

		tokens, refreshToken, _ := tokengen.TokenGenerator(*foundUser.Email, *foundUser.FistName, *foundUser.LastName, foundUser.UserID)
		defer cancel()

		tokengen.UpdateAllTokens(tokens, refreshToken, foundUser.UserID)
		c.JSON(http.StatusFound, foundUser)

	}
}

func ProductViewerAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var products dto.Product
		if err := c.BindJSON(&products); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		products.ProductID = primitive.NewObjectID()
		_, ins := ProductCollection.InsertOne(ctx, products)
		if ins != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Not created"})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, "Successfully added our Product Admin!!")

	}
}

func SearchProduct() gin.HandlerFunc {}

func SearchProductByQuery() gin.HandlerFunc {}
