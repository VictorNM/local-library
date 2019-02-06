package model

import (
	"database/sql"
	"log"

	database "../shared"
)

// Book DTO
type Book struct {
	ID       string
	Title    string
	AuthorID string
	Summary  string
	ISBN     string
}

// GetBooks return list of books in database
func GetBooks() []Book {
	db := database.GetDB()
	rows, err := db.Query(`SELECT * FROM books`)
	if err != nil {
		log.Fatal(err.Error())
	}

	var books []Book
	for rows.Next() {
		var id string
		var title string
		var authorID string
		var summary string
		var isbn string
		err = rows.Scan(&id, &title, &authorID, &summary, &isbn)
		book := Book{ID: id, Title: title, AuthorID: authorID, Summary: summary, ISBN: isbn}
		books = append(books, book)
	}
	return books
}

func parseBook(row *sql.Row) (*Book, error) {
	var book = &Book{}
	err := row.Scan(&book.ID, &book.Title, &book.AuthorID, &book.Summary, &book.ISBN)
	return book, err
}

func parseBooks(rows *sql.Rows) ([]*Book, error) {
	var books []*Book
	for rows.Next() {
		book := &Book{}
		err := rows.Scan(&book.ID, &book.Title, &book.AuthorID, &book.Summary, &book.ISBN)

		if err != nil {
			return nil, err
		}

		books = append(books, book)
	}
	return books, nil
}
