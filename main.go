package main

import (
	"log"
	"net/http"

	"./web/app/route"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/catalog/", http.StatusFound)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/catalog/", route.CatalogHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err.Error())
	}
}
