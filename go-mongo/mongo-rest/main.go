package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

type Person struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Firstname string             `json:"firstname,omitempty" bson:"firstname,omitempty"`
	Lastname  string             `json:"lastname,omitempty" bson:"lastname,omitempty"`
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var person Person
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("mongo-rest").Collection("person")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, err := collection.InsertOne(ctx, person)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(result)

}
func GetAllperson(w http.ResponseWriter, r *http.Request) {
	fmt.Println("check")
}
func GetPerson(w http.ResponseWriter, r *http.Request)    {}
func UpdatePerson(w http.ResponseWriter, r *http.Request) {}
func DeletePerson(w http.ResponseWriter, r *http.Request) {}

func main() {
	fmt.Println("main function")

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, _ = mongo.Connect(ctx, clientOptions)
	defer client.Disconnect(ctx)

	router := mux.NewRouter()

	router.HandleFunc("/person", CreatePerson).Methods("POST")
	router.HandleFunc("/person", GetAllperson).Methods("GET")
	router.HandleFunc("/person/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/person/{id}", UpdatePerson).Methods("PUT")
	router.HandleFunc("/person/{id}", DeletePerson).Methods("DELETE")

	http.ListenAndServe(":8080", router)

}
