package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreatePerson(w http.ResponseWriter, r *http.Request) {}
func GetAllperson(w http.ResponseWriter, r *http.Request) {}
func GetPerson(w http.ResponseWriter, r *http.Request)    {}
func UpdatePerson(w http.ResponseWriter, r *http.Request) {}
func DeletePerson(w http.ResponseWriter, r *http.Request) {}

func main() {
	fmt.Println("main function")

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	router := mux.NewRouter()

	router.HandleFunc("/person", CreatePerson).Methods("POST")
	router.HandleFunc("/person", GetAllperson).Methods("GET")
	router.HandleFunc("/person/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/person/{id}", UpdatePerson).Methods("PUT")
	router.HandleFunc("/person/{id}", DeletePerson).Methods("DELETE")

	http.ListenAndServe(":8080", router)

}
