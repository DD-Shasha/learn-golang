package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// GetItems returns all items
func GetItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Items)
}

// GetItem returns a single item by ID
func GetItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range Items {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Item{})
}

// CreateItem adds a new item
func CreateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var item Item
	_ = json.NewDecoder(r.Body).Decode(&item)
	Items = append(Items, item)
	json.NewEncoder(w).Encode(item)
}

// UpdateItem updates an existing item
func UpdateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range Items {
		if item.ID == params["id"] {
			Items = append(Items[:index], Items[index+1:]...)
			var item Item
			_ = json.NewDecoder(r.Body).Decode(&item)
			item.ID = params["id"]
			Items = append(Items, item)
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(Items)
}

// DeleteItem removes an item
func DeleteItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range Items {
		if item.ID == params["id"] {
			Items = append(Items[:index], Items[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(Items)
}
