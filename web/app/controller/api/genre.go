package api

import (
	"encoding/json"
	"net/http"

	"../../model"
)

// GetGenres handle /api/catalog/genres
func GetGenres(w http.ResponseWriter, r *http.Request) {
	genres, _ := model.GetGenres()
	json.NewEncoder(w).Encode(genres)
}
