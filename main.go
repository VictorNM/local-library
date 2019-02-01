package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"./web/app/route"
	database "./web/app/shared"
)

// db info
const (
	DatabaseUser     = "postgres"
	DatabasePassword = "admin"
	DatabaseName     = "local_library"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/catalog/", http.StatusFound)
}

func main() {
	// connect to db
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DatabaseUser, DatabasePassword, DatabaseName)
	database.Connect(dbinfo)
	defer database.Close()

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/catalog/", route.CatalogHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err.Error())
	}
}
