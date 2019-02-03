package route

import (
	"log"
	"net/http"

	"../controller"
)

// CatalogHandler route for "/catalog/"
func CatalogHandler(w http.ResponseWriter, r *http.Request) {
	catalog := r.URL.Path[len("/catalog/"):]
	log.Printf("Method: %s  \t|\tPath: /catalog/%s", r.Method, catalog)

	switch catalog {
	case "":
		http.Redirect(w, r, catalog+"books", http.StatusFound)
	case "books":
		controller.GetBooks(w, r)
	case "book/create":
		controller.CreateBook(w, r)
	case "book/update":
		controller.UpdateBook(w, r)
	default:
		http.Error(w, "NOT FOUND", http.StatusNotFound)
	}
}
