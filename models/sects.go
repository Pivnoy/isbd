package models

import "database/sql"

type SectsResponse struct {
	Sects []Sects `json:"sects"`
}

type Sects struct {
	Id   int
	Name string
}

func parseSects(rows *sql.Rows) (Sects, error) {
	c := Sects{}
	err := rows.Scan(&c.Id, &c.Name)
	if err != nil {
		panic("Error in parse Channel")
	}
	return c, nil
}

func CreateSectsCollection(rows *sql.Rows) ([]Sects, error) {
	var sectsCollection []Sects
	for rows.Next() {
		j, err := parseSects(rows)
		if err != nil {
			panic("Error in parsing array of Channels")
		}
		sectsCollection = append(sectsCollection, j)
	}
	return sectsCollection, nil
}
