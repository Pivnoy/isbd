package models

import "database/sql"

type JoblessResponse struct {
	Jobless   float32 `json:"jobless"`
	Jobless_c int     `json:"jobless_c"`
}

type Human struct {
	Id         int
	Name       string
	Weight     int
	Employment bool
	Salary     int
	Country_id string
}

func parseHuman(rows *sql.Rows) (Human, error) {
	c := Human{}
	err := rows.Scan(&c.Id, &c.Name, &c.Weight, &c.Employment, &c.Salary, &c.Country_id)
	if err != nil {
		panic("Error in parse Channel")
	}
	return c, nil
}

func CreateHumanCollection(rows *sql.Rows) ([]Human, error) {
	var joblessCollection []Human
	for rows.Next() {
		j, err := parseHuman(rows)
		if err != nil {
			panic("Error in parsing array of Channels")
		}
		joblessCollection = append(joblessCollection, j)
	}
	return joblessCollection, nil
}
