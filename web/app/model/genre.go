package model

import (
	"database/sql"
	"log"

	database "../shared"
)

// Genre DTO
type Genre struct {
	ID    string
	Name  string
	Books []*Book
}

// GetGenres query all genres
func GetGenres() ([]*Genre, error) {
	log.Println("Querying...")
	rows, err := findAll("genres")

	log.Println("Query finish")
	if err != nil {
		log.Printf("Error in GetGenres: %s", err.Error())
		return nil, err
	}

	log.Println("Parsing")
	genres, err := parseGenres(rows)
	log.Println("Parse finish")

	if err != nil {
		log.Printf("Error in GetGenres: %s", err.Error())
		return nil, err
	}

	return genres, nil
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
func FindGenreByID(id string) (Genre, error) {
	genre := Genre{}
	row := findByID("genres", id)
	err := row.Scan(&genre.ID, &genre.Name)

	if err != nil {
		return Genre{}, err
	}

	genre.Books, err = findBooksByGenreID(id)

	return genre, err
}

func findBooksByGenreID(id string) ([]*Book, error) {
	condition := make(map[string]string)
	condition["genre_id"] = id
	rows, err := find("books_genres", condition)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	var books []*Book
	for rows.Next() {
		var bookID string
		var genreID string // use to ignore value genre_id
		err := rows.Scan(&bookID, &genreID)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}
		row := findByID("books", bookID)
		book, err := parseBook(row)
		books = append(books, book)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}
	}
	return books, err
}

func parseGenres(rows *sql.Rows) ([]*Genre, error) {
	var genres []*Genre
	for rows.Next() {
		genre := &Genre{}
		err := rows.Scan(&genre.ID, &genre.Name)
		if err != nil {
			return nil, err
		}
		genre.Books, err = findBooksByGenreID(genre.ID)
		if err != nil {
			return nil, err
		}
		genres = append(genres, genre)
	}
	return genres, nil
}

func parseGenre(row *sql.Row) (*Genre, error) {
	genre := &Genre{}
	err := row.Scan(&genre.ID, &genre.Name)
	if err != nil {
		return nil, err
	}
	genre.Books, err = findBooksByGenreID(genre.ID)
	if err != nil {
		return nil, err
	}
	return genre, err
}
