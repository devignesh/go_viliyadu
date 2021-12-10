package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func CreatePerson(w http.ResponseWriter, r *http.Request) {}
func GetAllperson(w http.ResponseWriter, r *http.Request) {}
func GetPerson(w http.ResponseWriter, r *http.Request)    {}
func UpdatePerson(w http.ResponseWriter, r *http.Request) {}
func DeletePerson(w http.ResponseWriter, r *http.Request) {}

func main() {
	fmt.Println("main function")

	router := mux.NewRouter()

	router.HandleFunc("/person", CreatePerson).Methods("POST")
	router.HandleFunc("/person", GetAllperson).Methods("GET")
	router.HandleFunc("/person/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/person/{id}", UpdatePerson).Methods("PUT")
	router.HandleFunc("/person/{id}", DeletePerson).Methods("DELETE")

	http.ListenAndServe(":8080", router)

}
