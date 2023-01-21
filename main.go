package main

import (
	"github.com/folklore13/golang-rest-api/controllers/productcontroller"
	"github.com/folklore13/golang-rest-api/models"
	"net/http"
	"github.com/gorilla/mux"
)

func main(){
	models.ConnectToDatabase()

	r := mux.NewRouter()

	r.HandleFunc("/products", productcontroller.Index).Methods("GET")
	r.HandleFunc("/product/{id}", productcontroller.Index).Methods("GET")
	r.HandleFunc("/product", productcontroller.Index).Methods("POST")
	r.HandleFunc("/product/{id}", productcontroller.Index).Methods("PUT")
	r.HandleFunc("/product", productcontroller.Index).Methods("DELETE")

	http.ListenAndServe(":8080", r)
}