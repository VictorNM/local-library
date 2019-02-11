package api

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"../../model"
	"github.com/gorilla/mux"
)

// GetGenres handle /api/catalog/genres
func GetGenres(w http.ResponseWriter, r *http.Request) {
	defer handleError(w, r)
	genres := model.FindAllGenres()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(genres)
}

// GetGenre handle /api/catalog/genre/{id}
func GetGenre(w http.ResponseWriter, r *http.Request) {
	defer handleError(w, r)
	id := mux.Vars(r)["id"]
	genre := model.FindGenreByID(id)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(genre)
}

// CreateGenre handle /api/catalog/genre/create
func CreateGenre(w http.ResponseWriter, r *http.Request) {
	defer handleError(w, r)
	genre := parseRequestToGenre(r)
	id := model.InsertGenre(genre)
	genre.ID = id
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(genre)
}

// UpdateGenre update genre
func UpdateGenre(w http.ResponseWriter, r *http.Request) {
	defer handleError(w, r)
	genre := parseRequestToGenre(r)
	genre.ID = mux.Vars(r)["id"]
	model.UpdateGenre(genre)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(genre)
}

// DeleteGenre delete genre
func DeleteGenre(w http.ResponseWriter, r *http.Request) {
	defer handleError(w, r)
	id := mux.Vars(r)["id"]
	model.DeleteGenre(id)
	w.WriteHeader(http.StatusNoContent)
}

func parseRequestToGenre(r *http.Request) model.Genre {
	var genre model.Genre
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	panicIfError(err)
	r.Body.Close()
	err = json.Unmarshal(body, &genre)
	panicIfError(err)
	return genre
}

func handleError(w http.ResponseWriter, r *http.Request) {
	if recover := recover(); recover != nil {
		log.Printf(`Recover due to error: "%s" from "%s"`, recover, r)

		switch x := recover.(type) {
		case string:
			json.NewEncoder(w).Encode(x)
		case error:
			json.NewEncoder(w).Encode(x.Error())
		default:
			json.NewEncoder(w).Encode(errors.New("unknown error"))
		}

	}
}

func panicIfError(err error) {
	if err != nil {
		panic(err)
	}
}
