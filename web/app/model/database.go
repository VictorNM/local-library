package model

import (
	"database/sql"
	"fmt"

	database "../shared"
)

func findAll(table string) (*sql.Rows, error) {
	db := database.GetDB()
	rows, err := db.Query(fmt.Sprintf("SELECT * FROM %s;", table))
	return rows, err
}

func findByID(table string, id string) *sql.Row {
	db := database.GetDB()
	idCol := table[:len(table)-1] + "_id"
	queryString := fmt.Sprintf("SELECT * FROM %s WHERE %s = %s;", table, idCol, id)
	row := db.QueryRow(queryString)
	return row
}

// Find query rows that match condition
func find(table string, condition map[string]string) (*sql.Rows, error) {
	db := database.GetDB()
	var k, v string
	for k, v = range condition {
		break
	}
	queryString := fmt.Sprintf("SELECT * FROM %s WHERE %s = %s;", table, k, v)
	rows, err := db.Query(queryString)
	return rows, err
}
