package route

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route define route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes array of Route
type Routes []Route

// NewRouter create router
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	setAPIRouter(router)

	return router
}
