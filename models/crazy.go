package models

import "database/sql"

type CrazyResponse struct {
	Crazy        float32 `json:"crazy"`
	CrazyC       int     `json:"crazy_c"`
	PeopleAmount int     `json:"people_amount"`
}

type Crazy struct {
	Id      int
	HumanId string
}

func parseCrazy(rows *sql.Rows) (Crazy, error) {
	c := Crazy{}
	err := rows.Scan(&c.Id, &c.HumanId)
	if err != nil {
		panic("Error in parse Channel")
	}
	return c, nil
}

func CreateCrazyCollection(rows *sql.Rows) ([]Crazy, error) {
	var crazyCollection []Crazy
	for rows.Next() {
		j, err := parseCrazy(rows)
		if err != nil {
			panic("Error in parsing array of Channels")
		}
		crazyCollection = append(crazyCollection, j)
	}
	return crazyCollection, nil
}
