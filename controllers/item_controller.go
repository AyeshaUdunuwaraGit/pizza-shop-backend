package controllers

import (
	"encoding/json"
	"net/http"
	"pizza-shop-backend/config"
	"pizza-shop-backend/models"
)

func GetAllItems(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query("SELECT id, name, category, price FROM items")
	if err != nil {
		http.Error(w, "Failed to query items", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var items []models.Item
	for rows.Next() {
		var item models.Item
		err := rows.Scan(&item.ID, &item.Name, &item.Category, &item.Price)
		if err != nil {
			http.Error(w, "Error scanning item", http.StatusInternalServerError)
			return
		}
		items = append(items, item)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}
func CreateItem(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST is allowed", http.StatusMethodNotAllowed)
		return
	}

	var item models.Item
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	query := "INSERT INTO items (name, category, price) VALUES ($1, $2, $3)"
	_, err = config.DB.Exec(query, item.Name, item.Category, item.Price)
	if err != nil {
		http.Error(w, "Failed to insert item", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Item added successfully"))
}
