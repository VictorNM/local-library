package route

import (
	"log"
	"net/http"
	"regexp"

	"../controller"
)

const (
	pathCatalog = "/catalog/"
	// subpaths
	pathGenreList = "genres"
	pathGenreView = "^genre/(.+)$"
)

// CatalogHandler route for "/catalog/"
func CatalogHandler(w http.ResponseWriter, r *http.Request) {
	subpath := r.URL.Path[len(pathCatalog):]
	log.Printf("Method: %s  \t|\tPath: /catalog/%s", r.Method, subpath)

	genreView := regexp.MustCompile(pathGenreView)

	if subpath == pathGenreList {
		controller.GetGenres(w, r)
	} else if genreView.MatchString(subpath) {
		controller.GetGenre(w, r)
	} else {
		http.Error(w, "NOT FOUND", http.StatusNotFound)
	}
}
