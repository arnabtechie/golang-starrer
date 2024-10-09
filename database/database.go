package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB(connStr string) {
	var err error

	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	DB.SetMaxOpenConns(25)
	DB.SetMaxIdleConns(25)
	DB.SetConnMaxLifetime(0)

	if err := DB.Ping(); err != nil {
		log.Fatalf("Unable to ping the database: %v\n", err)
	}

	fmt.Println("Connected to PostgreSQL database!")
}

func CloseDB() {
	if err := DB.Close(); err != nil {
		log.Fatalf("Error closing the database: %v\n", err)
	}
	fmt.Println("Database connection closed!")
}
