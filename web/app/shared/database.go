package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // database library
)

var database *sql.DB

// Connect to db
func Connect(dataSourceName string) *sql.DB {
	db, err := sql.Open("postgres", dataSourceName)

	if err != nil {
		log.Fatal(err.Error())
	}

	return db
}

// Close db
func Close() {
	database.Close()
}

// GetDB return database instance
func GetDB() *sql.DB {
	return database
}
