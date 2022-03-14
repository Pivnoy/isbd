package models

import "database/sql"

type Jobless struct {
	jobless   float32 `json:"jobless"`
	jobless_c int     `json:"jobless_c"`
}

func parseJobless(rows *sql.Rows) (Jobless, error) {
	c := Jobless{}
	err := rows.Scan(&c.jobless, &c.jobless_c)
	if err != nil {
		panic("Error in parse Channel")
	}
	return c, nil
}

func CreateJoblessCollection(rows *sql.Rows) ([]Jobless, error) {
	var joblessCollection []Jobless
	for rows.Next() {
		j, err := parseJobless(rows)
		if err != nil {
			panic("Error in parsing array of Channels")
		}
		joblessCollection = append(joblessCollection, j)
	}
	return joblessCollection, nil
}
