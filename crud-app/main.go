package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize the router
	r := mux.NewRouter()

	// Seed some initial data
	Items = append(Items, Item{ID: "1", Name: "Item 1", Description: "First Item"})

	// Route handles & endpoints
	r.HandleFunc("/items", GetItems).Methods("GET")
	r.HandleFunc("/items/{id}", GetItem).Methods("GET")
	r.HandleFunc("/items", CreateItem).Methods("POST")
	r.HandleFunc("/items/{id}", UpdateItem).Methods("PUT")
	r.HandleFunc("/items/{id}", DeleteItem).Methods("DELETE")

	// Start server
	fmt.Println("Server is starting on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", r))
}
