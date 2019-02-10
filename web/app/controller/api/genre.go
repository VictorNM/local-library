package api

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"../../model"
	"github.com/gorilla/mux"
)

// GetGenres handle /api/catalog/genres
func GetGenres(w http.ResponseWriter, r *http.Request) {
	genres, err := model.FindAllGenres()
	if err != nil {
		log.Println(err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(genres)
}

// GetGenre handle /api/catalog/genre/{id}
func GetGenre(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	genre, err := model.FindGenreByID(id)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(genre)
}

// CreateGenre handle /api/catalog/genre/create
func CreateGenre(w http.ResponseWriter, r *http.Request) {
	var genre model.Genre
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		log.Println(err.Error())
		return
	}
	r.Body.Close()
	err = json.Unmarshal(body, &genre)
	if err != nil {
		log.Println(err.Error())
		return
	}
	id, err := model.InsertGenre(genre)
	if err != nil {
		log.Println(err.Error())
		return
	}
	genre.ID = id
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(genre)
}

// UpdateGenre update genre
func UpdateGenre(w http.ResponseWriter, r *http.Request) {
	genre, err := parseRequestToGenre(r)
	if err != nil {
		log.Println(err.Error())
		return
	}
	genre.ID = mux.Vars(r)["id"]
	_, err = model.UpdateGenre(genre)
	if err != nil {
		log.Println(err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(genre)
}

func parseRequestToGenre(r *http.Request) (model.Genre, error) {
	var genre model.Genre
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		return genre, err
	}
	r.Body.Close()
	err = json.Unmarshal(body, &genre)
	if err != nil {
		return genre, err
	}
	return genre, nil
}
