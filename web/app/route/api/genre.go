package api

import (
	"victornm/library/web/app/controller/api"
	"victornm/library/web/app/route/route"
)

func getGenreRoutes() route.Routes {
	return route.Routes{
		route.Route{
			Name:        "apiGetGenres",
			Method:      "GET",
			Pattern:     "/catalog/genres",
			HandlerFunc: api.GetGenres,
		},
		route.Route{
			Name:        "apiGetGenre",
			Method:      "GET",
			Pattern:     "/catalog/genre/{id:[1-9][0-9]*}",
			HandlerFunc: api.GetGenre,
		},
		route.Route{
			Name:        "apiCreateGenre",
			Method:      "POST",
			Pattern:     "/catalog/genre/create",
			HandlerFunc: api.CreateGenre,
		},
		route.Route{
			Name:        "apiUpdateGenre",
			Method:      "PUT",
			Pattern:     "/catalog/genre/{id:[1-9][0-9]*}",
			HandlerFunc: api.UpdateGenre,
		},
		route.Route{
			Name:        "apiDeleteGenre",
			Method:      "DELETE",
			Pattern:     "/catalog/genre/{id:[1-9][0-9]*}",
			HandlerFunc: api.DeleteGenre,
		},
	}
}
