package route

import (
	"./api"
	"github.com/gorilla/mux"
)

// NewRouter create router
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	api.SetAPIRouter(router)

	return router
}
