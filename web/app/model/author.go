package model

import (
	"time"
)

// Author DAO
type Author struct {
	ID          string
	FirstName   string
	LastName    string
	DateOfBirth time.Time
	DateOfDeath time.Time
}

// Name virtual property for author's name
func (author *Author) Name() string {
	return author.FirstName + " " + author.LastName
}
