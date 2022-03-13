package models

import "database/sql"

type Channel struct {
	Id     uint64
	Number int
	Name   string
	Theme  string
}

func ParseChannel(rows *sql.Rows) Channel {
	c := Channel{}
	err := rows.Scan(&c.Id, &c.Number, &c.Name, &c.Theme)
	if err != nil {
		panic("Error in parse Channel")
	}
	return c
}
