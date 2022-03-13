package server

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" // add this
)

var DbInstance *sql.DB

func Init(initArgs []string) {
	var err error
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", initArgs[0], initArgs[1], initArgs[2], initArgs[3], initArgs[4])
	DbInstance, err = sql.Open("postgres", connStr)
	if err != nil {
		panic("Error in connection to PostgreSQL")
	}
	err = DbInstance.Ping()
	if err != nil {
		panic(err)
	}
}
