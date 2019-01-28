package route

import (
	"log"
	"net/http"

	"../controller"
)

// CatalogHandler route for "/catalog/"
func CatalogHandler(w http.ResponseWriter, r *http.Request) {
	catalog := r.URL.Path[len("/catalog/"):]
	log.Printf("Path: /catalog/%s", catalog)

	switch catalog {
	case "":
		http.Redirect(w, r, catalog+"books/", http.StatusFound)
	case "books/":
		controller.GetBooks(w, r)
	case "book/create/":
		controller.CreateBook(w, r)
	default:
		http.Error(w, "NOT FOUND", http.StatusNotFound)
	}
}
