package controller

import (
	"fmt"
	"net/http"
)

// Books GET books list
func GetBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "NOT IMPLEMENTED BOOKS")
}

// CreateBook handle create book
func CreateBook(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	switch method {
	case "GET":
		getCreateBook(w, r)
	case "POST":
		postCreateBook(w, r)
	}
}

func getCreateBook(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "NOT IMPLEMENTED BOOK_CREATE_GET")
}

func postCreateBook(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "NOT IMPLEMENTED BOOK_CREATE_POST")
}
