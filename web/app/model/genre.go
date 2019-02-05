package model

import (
	"log"

	database "../shared"
)

// Genre DTO
type Genre struct {
	ID   string
	Name string
}

// GetGenres query all genres
func GetGenres() []Genre {
	db := database.GetDB()
	rows, err := db.Query(`SELECT * FROM genres`)
	if err != nil {
		log.Fatal(err.Error())
	}

	var genres []Genre
	for rows.Next() {
		var id string
		var name string
		err = rows.Scan(&id, &name)
		genre := Genre{ID: id, Name: name}
		genres = append(genres, genre)
	}
	return genres
}

// InsertGenre insert new genre to db
func InsertGenre(genre Genre) string {
	db := database.GetDB()

	var lastInsertID string
	err := db.QueryRow(`INSERT INTO genres (name) VALUES ($2) RETURNING genre_id;`, genre.Name).Scan(&lastInsertID)

	if err != nil {
		log.Fatal(err.Error())
	}

	return lastInsertID
}

// FindGenreByID find genre by ID
func FindGenreByID(id string) Genre {
	db := database.GetDB()
	genre := Genre{}
	row := db.QueryRow(`SELECT * FROM genres WHERE genre_id=$1`, id)
	err := row.Scan(&genre.ID, &genre.Name)
	if err != nil {
		log.Fatal(err.Error())
	}
	return genre
}
