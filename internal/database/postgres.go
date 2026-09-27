package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // PostgreSQL driver (is a must-have for PostgreSQL)

	"github.com/kholiklutfi29/nourish-dispatch/internal/config"
)

func ConnectPostgres() *sql.DB {
	databaseURL := config.GetDatabaseURL()

	db, err := sql.Open("postgres", databaseURL)

	if err != nil {
		log.Fatal("Failed to open database:", err)
	}

	err = db.Ping()

	if err != nil {
		log.Fatal("Failed to connect to PostgreSQL:", err)
	}

	fmt.Println("Successfully connected to PostgreSQL database")

	return db
}