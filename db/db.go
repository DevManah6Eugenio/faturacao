package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	USER     = "postgres"
	PASSWORD = "123"
	DATABASE = "faturacao"
	HOST     = "127.0.0.1"
	PORT     = 5432
	SSLMODE  = "disable"
)

func Connection() (*sql.DB) { 

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		HOST, PORT, USER, PASSWORD, DATABASE, SSLMODE)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err.Error())
	}

	return db
}
