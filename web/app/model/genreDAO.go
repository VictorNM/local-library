package model

import (
	"database/sql"
	"fmt"

	"../database"
)

const (
	genreTable    = "genres"
	genreIDColumn = "genre_id"
)

// FindAllGenres query all genre from db
func FindAllGenres() (Genres, error) {
	db := database.GetDB()
	queryString := fmt.Sprintf("SELECT * FROM %s", genreTable)
	rows, err := db.Query(queryString)
	if err != nil {
		return nil, err
	}
	return parseGenres(rows)
}

// FindGenreByID query genre by ID
func FindGenreByID(id string) (Genre, error) {
	db := database.GetDB()
	queryString := fmt.Sprintf("SELECT * FROM %s WHERE %s=%s", genreTable, genreIDColumn, id)
	row := db.QueryRow(queryString)
	return parseGenre(row)
}

// InsertGenre insert genre to db
func InsertGenre(genre *Genre) (string, error) {
	db := database.GetDB()
	queryString := fmt.Sprintf(`INSERT INTO genres (name) VALUES ('%s') RETURNING genre_id`, genre.Name)
	var insertID string
	err := db.QueryRow(queryString).Scan(&insertID)
	return insertID, err
}

func parseGenres(rows *sql.Rows) (Genres, error) {
	var genres Genres
	for rows.Next() {
		genre := &Genre{}
		err := rows.Scan(&genre.ID, &genre.Name)
		if err != nil {
			return nil, err
		}
		genres = append(genres, genre)
	}
	return genres, nil
}

func parseGenre(row *sql.Row) (Genre, error) {
	var genre = Genre{}
	err := row.Scan(&genre.ID, &genre.Name)
	return genre, err
}
