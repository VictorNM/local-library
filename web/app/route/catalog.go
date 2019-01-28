package route

import (
	"fmt"
	"log"
	"net/http"
)

// CatalogHandler route for "/catalog/"
func CatalogHandler(w http.ResponseWriter, r *http.Request) {
	catalog := r.URL.Path[len("/catalog/"):]
	log.Println(catalog)

	switch catalog {
	case "":
		http.Redirect(w, r, catalog+"book", http.StatusFound)
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
