package controller

import (
	"html/template"
	"log"
	"net/http"

	"../model"
)

// GetGenres handle request /catalog/genres
func GetGenres(w http.ResponseWriter, r *http.Request) {
	genres := model.GetGenres()
	t, err := template.ParseFiles("./web/template/genre-list.html")
	if err != nil {
		log.Fatal(err.Error())
	}
	t.Execute(w, genres)
}

// CreateGenre handle /catalog/genre/create
func CreateGenre(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	switch method {
	case http.MethodGet:
		getCreateGenre(w, r)
	case http.MethodPost:
		postCreateGenre(w, r)
	}
}

func getCreateGenre(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./web/template/create-genre.html")
	if err != nil {
		log.Fatal(err.Error())
	}
	t.Execute(w, nil)
}

func postCreateGenre(w http.ResponseWriter, r *http.Request) {
	genre := model.Genre{ID: "gen010", Name: r.FormValue("name")}
	model.InsertGenre(genre)
}
