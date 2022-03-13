package server

import (
	"database/sql"
	_ "github.com/lib/pq" // add this
)

func CreateDBInstance() *sql.DB {
	db, err := sql.Open("postgres", "postgresql://postgres:lehavnuk@postgres")
	if err != nil {
		panic("Error in connection to PostgreSQL")
	}
	return db
}
