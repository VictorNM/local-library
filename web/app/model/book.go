package model

import (
	"log"

	database "../shared"
)

// Book DAO
type Book struct {
	Title   string
	Summary string
}

// GetBooks return list of books in database
func GetBooks() []Book {
	db := database.GetDB()
	rows, err := db.Query(`SELECT * FROM book`)
	if err != nil {
		log.Fatal(err.Error())
	}

	var books []Book
	for rows.Next() {
		var title string
		var summary string
		err = rows.Scan(&title, &summary)
		book := Book{Title: title, Summary: summary}
		books = append(books, book)
	}
	return books
}
