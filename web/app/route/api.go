package route

import (
	"../controller/api"
	"github.com/gorilla/mux"
)

var apiRoutes = Routes{
	Route{
		Name:        "apiGetGenres",
		Method:      "GET",
		Pattern:     "/catalog/genres",
		HandlerFunc: api.GetGenres,
	},
	Route{
		Name:        "apiGetGenre",
		Method:      "GET",
		Pattern:     "/catalog/genre/{id:[1-9][0-9]*}",
		HandlerFunc: api.GetGenre,
	},
	Route{
		Name:        "apiCreateGenre",
		Method:      "POST",
		Pattern:     "/catalog/genre/create",
		HandlerFunc: api.CreateGenre,
	},
}

func setAPIRouter(router *mux.Router) {
	s := router.PathPrefix("/api").Subrouter()
	for _, route := range apiRoutes {
		s.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
}
