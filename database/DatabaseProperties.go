package database

import "os"

const (
	DBHOST  = os.Getenv("DBHOST")
	DBPORT  = os.Getenv("DBPORT")
	DBUSER  = os.Getenv("DBUSER")
	DBPASS  = os.Getenv("DBPASS")
	DBNAME  = os.Getenv("DBNAME")
	SSLMODE = os.Getenv("SSLMODE")
)
