package main

import (
	"log"
	"net/http"
	"pizza-shop-backend/config"
	"pizza-shop-backend/routes"
)

func main() {
	config.ConnectDB()
	routes.RegisterRoutes()
	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
