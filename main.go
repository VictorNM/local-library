package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"victornm/library/web/app/database"
	"victornm/library/web/app/route"
)

// db info
const (
	DatabaseUser     = "postgres"
	DatabasePassword = "admin"
	DatabaseName     = "local_library"
)

func main() {
	// DATABASE
	log.Println("Connecting database...")
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DatabaseUser, DatabasePassword, DatabaseName)
	database.Connect(dbinfo)
	defer database.Close()

	// ROUTING
	router := route.NewRouter()

	// START SERVER
	port := 8080
	log.Printf("Server started at port: %d\n", port)
	log.Fatal(http.ListenAndServe(":8080", router))
}
