package main

// Item represents a single item in our CRUD application
type Item struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Items is a slice to store multiple items
var Items []Item
