package api

import (
	"../route"
	"github.com/gorilla/mux"
)

// SetAPIRouter set router for /api/...
func SetAPIRouter(router *mux.Router) {
	var routes route.Routes
	routes = append(routes, getGenreRoutes()...)
	s := router.PathPrefix("/api").Subrouter()
	for _, route := range routes {
		s.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
}
