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
	log.Println("Connecting database...")
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DatabaseUser, DatabasePassword, DatabaseName)
	database.Connect(dbinfo)
	defer database.Close()

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/catalog/", route.CatalogHandler)

	port := 8080
	log.Printf("Server started at port: %d\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatal(err.Error())
	}

}
