package config

import (
	"database/sql"
	"log"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
	connStr := "host=localhost port=5432 user=postgres password=@1234 dbname=pizza_shop sslmode=disable"
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	err = DB.Ping()
	if err != nil {
		log.Fatal("Cannot reach database:", err)
	}
	log.Println("Connected to PostgreSQL database!")
}
