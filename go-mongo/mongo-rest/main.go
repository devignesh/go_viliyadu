package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//mongo client setting
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
	w.Header().Set("content-type", "application/json")
	var person []Person
	collection := client.Database("mongo-rest").Collection("person")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	result, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer result.Close(ctx)

	for result.Next(ctx) {
		var person1 Person
		result.Decode(&person1)
		person = append(person, person1)
	}

	json.NewEncoder(w).Encode(person)

}

func GetPerson(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var person Person
	collection := client.Database("mongo-rest").Collection("person")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err := collection.FindOne(ctx, Person{ID: id}).Decode(&person)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(person)
}
func UpdatePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var person Person

	_ = json.NewDecoder(r.Body).Decode(&person)

	collection := client.Database("mongo-rest").Collection("person")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	update := bson.D{
		{
			"$set", bson.D{
				{"firstname", person.Firstname},
				{"lastname", person.Lastname},
			},
		},
	}
	err := collection.FindOneAndUpdate(ctx, Person{ID: id}, update).Decode(&person)
	if err != nil {
		log.Fatal(err)
	}
	person.ID = id
	json.NewEncoder(w).Encode(person)

}
func DeletePerson(w http.ResponseWriter, r *http.Request) {
	fmt.Println("delete")
	w.Header().Set("content-type", "application/json")
	var params = mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	filter := bson.M{"_id": id}
	collection := client.Database("mongo-rest").Collection("person")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	deleteperson, err := collection.DeleteOne(ctx, filter)

	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(deleteperson)

}

func main() {
	fmt.Println("main function")

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, _ = mongo.Connect(ctx, clientOptions)
	defer client.Disconnect(ctx)

	router := mux.NewRouter()

	router.HandleFunc("/person", CreatePerson).Methods("POST")
	router.HandleFunc("/persons", GetAllperson).Methods("GET")
	router.HandleFunc("/person/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/person/{id}", UpdatePerson).Methods("PUT")
	router.HandleFunc("/person/{id}", DeletePerson).Methods("DELETE")

	http.ListenAndServe(":8080", router)

}
