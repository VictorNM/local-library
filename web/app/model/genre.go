package model

// Genre DTO
type Genre struct {
	ID   string `json: "id"`
	Name string `json: "name"`
}

// Genres represent array of genres
type Genres []*Genre
