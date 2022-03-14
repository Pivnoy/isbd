package models

import "database/sql"

type PunitiveResponse struct {
	Structures []Punitive `json:"structures"`
}

type Punitive struct {
	Id   int
	Name string
}

func parsePunitive(rows *sql.Rows) (Punitive, error) {
	c := Punitive{}
	err := rows.Scan(&c.Id, &c.Name)
	if err != nil {
		panic("Error in parse Channel")
	}
	return c, nil
}

func CreatePunitiveCollection(rows *sql.Rows) ([]Punitive, error) {
	var punitiveCollection []Punitive
	for rows.Next() {
		j, err := parsePunitive(rows)
		if err != nil {
			panic("Error in parsing array of Channels")
		}
		punitiveCollection = append(punitiveCollection, j)
	}
	return punitiveCollection, nil
}
