package models

import "database/sql"

type Jobless struct {
	jobless   float32 `json:"jobless"`
	jobless_c int     `json:"jobless_c"`
}

func ParseJobless(rows *sql.Rows) Jobless {
	c := Jobless{}
	err := rows.Scan(&c.jobless, &c.jobless_c)
	if err != nil {
		panic("Error in parse Channel")
	}
	return c
}
