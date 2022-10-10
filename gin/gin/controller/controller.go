package controller

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Todo struct {
	ID        primitive.ObjectID `json:"id" bson:"_id, omitempty"`
	Title     string             `json:"title"`
	Body      string             `json:"body"`
	Completed string             `json:"completed"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}

var collection *mongo.Collection

// func Connection() *mongo.Client {

// 	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	err = client.Connect(ctx)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	// defer client.Disconnect(ctx)

// 	err = client.Ping(context.Background(), readpref.Primary())
// 	if err != nil {
// 		log.Fatal("Couldn't connect to the database", err)
// 	} else {
// 		log.Println("Connected!")
// 	}
// 	db := client.Database("mongo_gin")
// 	TodoCollection(db)
// 	return client

// }

func TodoCollection(c *mongo.Database) {
	collection = c.Collection("todos")
}

func CreateTodopost(c *gin.Context) {
	
	var ctx, _ = context.WithTimeout(context.Background(), 100*time.Second)
	var todo Todo
	c.BindJSON(&todo)
	title := todo.Title
	body := todo.Body
	completed := todo.Completed
	id := primitive.NewObjectID()

	newTodo := Todo{
		ID:        id,
		Title:     title,
		Body:      body,
		Completed: completed,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err := collection.InsertOne(ctx, newTodo)

	if err != nil {
		log.Printf("Error while inserting new todo into db, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Todo created Successfully",
	})
	return
}

func GetAllTodo(c *gin.Context) {
	
	todos := []Todo{}
	cursor, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		log.Printf("Error while getting all todos, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	//iterate the document
	for cursor.Next(context.TODO()) {
		var todo Todo
		cursor.Decode(&todo)
		todos = append(todos, todo)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "All Todos",
		"data":    todos,
	})
	return
}

func GetTodo(c *gin.Context) {
	
	id := c.Param("id")
	todo := Todo{}
	
	err := collection.FindOne(context.TODO(), bson.M{"id": id}).Decode(&todo)
	if err != nil {
		log.Printf("Error while getting a single todo, Reason: %v\n", err)
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Todo not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Single Todo",
		"data":    todo,
	})
	return
}

func UpdateTodo(c *gin.Context) {

	id := c.Param("id")
	var todo Todo
	c.BindJSON(&todo)
	completed := todo.Completed
	newData := bson.M{
		"$set": bson.M{
			"completed":  completed,
			"updated_at": time.Now(),
		},
	}
	_, err := collection.UpdateOne(context.TODO(), bson.M{"id": id}, newData)
	if err != nil {
		log.Printf("Error, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "Something went wrong",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Todo Edited Successfully",
	})
	return
}

func DeleteTodo(c *gin.Context) {
	
	id := c.Param("id")
	_, err := collection.DeleteOne(context.TODO(), bson.M{"id": id})
	if err != nil {
		log.Printf("Error while deleting a single todo, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Todo deleted successfully",
	})
	return
}
