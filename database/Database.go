package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func Connection() *sql.DB {

	db, err := sql.Open("postgres", driverConfig())

	if err != nil {
		panic(err.Error())
	}

	return db
}
