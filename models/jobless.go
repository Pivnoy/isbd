package models

import "database/sql"

type JoblessResponse struct {
	Jobless   float32 `json:"jobless"`
	Jobless_c int     `json:"jobless_c"`
}

type Jobless struct {
	Id         int
	Name       string
	Weight     int
	Employment bool
	Salary     int
	Country_id string
}

func parseJobless(rows *sql.Rows) (Jobless, error) {
	c := Jobless{}
	err := rows.Scan(&c.Id, &c.Name, &c.Weight, &c.Employment, &c.Salary, &c.Country_id)
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
