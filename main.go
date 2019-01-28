package main

import (
	"fmt"
	"log"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello")
}

func catalogHandler(w http.ResponseWriter, r *http.Request) {
	catalog := r.URL.Path[len("/catalog/"):]
	fmt.Println(catalog)
	switch catalog {
	case "book":
		bookHandler(w, r)
	case "author":
		authorHandler(w, r)
	case "genre":
		genreHandler(w, r)
	case "bookinstance":
		bookInstanceHandler(w, r)
	default:
		http.Error(w, "NOT FOUND", http.StatusNotFound)
	}
}

func bookHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "NOT IMPLEMENTED")
}

func authorHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "NOT IMPLEMENTED")
}

func genreHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "NOT IMPLEMENTED")
}

func bookInstanceHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "NOT IMPLEMENT")
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/catalog/", catalogHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err.Error())
	}
}
