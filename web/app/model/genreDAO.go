package model

import (
	"database/sql"
	"fmt"
	"log"

	"../database"
)

const (
	genreTable    = "genres"
	genreIDColumn = "genre_id"
)

// FindAllGenres query all genre from db
func FindAllGenres() Genres {
	db := database.GetDB()
	queryString := fmt.Sprintf("SELECT * FROM %s", genreTable)
	rows, err := db.Query(queryString)
	if err != nil {
		panic(err)
	}
	genres, err := parseGenres(rows)
	if err != nil {
		panic(err)
	}
	return genres
}

// FindGenreByID query genre by ID
func FindGenreByID(id string) Genre {
	db := database.GetDB()
	queryString := fmt.Sprintf("SELECT * FROM %s WHERE %s=%s", genreTable, genreIDColumn, id)
	row := db.QueryRow(queryString)
	genre, err := parseGenre(row)
	panicIfError(err)
	return genre
}

// InsertGenre insert genre to db
func InsertGenre(genre Genre) (string, error) {
	db := database.GetDB()
	queryString := fmt.Sprintf(`INSERT INTO genres (name) VALUES ('%s') RETURNING genre_id`, genre.Name)
	var insertID string
	err := db.QueryRow(queryString).Scan(&insertID)
	return insertID, err
}

// UpdateGenre update genre in db
func UpdateGenre(genre Genre) error {
	db := database.GetDB()
	log.Println("Updating...")
	stmt, err := db.Prepare(`UPDATE genres SET name=$1 WHERE genre_id=$2 RETURNING genre_id`)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(genre.Name, genre.ID)
	if err != nil {
		return err
	}
	affect, err := res.RowsAffected()
	if err != nil {
		return err
	}
	fmt.Println(affect, "rows change")
	return nil
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

func panicIfError(err error) {
	if err != nil {
		panic(err)
	}
}
