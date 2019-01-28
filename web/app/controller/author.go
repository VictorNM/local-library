package controller

import (
	"fmt"
	"net/http"
)

// Authors GET authors list
func Authors(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "NOT IMPLEMENTED AUTHORS/")
}
