package route

import (
	"net/http"
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
