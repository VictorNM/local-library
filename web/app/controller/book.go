package controller

import (
	"fmt"
	"html/template"
	"net/http"
)

// GetBooks books list
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
	t, _ := template.ParseFiles("./web/template/create-book.html")
	t.Execute(w, nil)
}

func postCreateBook(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "NOT IMPLEMENTED BOOK_CREATE_POST")
}

// UpdateBook handler /book/update
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	switch method {
	case "GET":
		getUpdateBook(w, r)
	case "POST":
		postUpdateBook(w, r)
	}
}

func getUpdateBook(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "NOT IMPLEMENTED GET_UPDATE_BOOK")
}

func postUpdateBook(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "NOT IMPLEMENTED POST_UPDATE_BOOK")
}
