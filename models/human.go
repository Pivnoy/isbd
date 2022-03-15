package models

import "database/sql"

type HumanDateResponse struct {
	Humans []HumanDat `json:"humans"`
}

type HumanDat struct {
	Id          int
	Name        string
	CountryName string
}

func parseHumanDatResponse(rows *sql.Rows) (HumanDat, error) {
	c := HumanDat{}
	err := rows.Scan(&c.Id, &c.Name, &c.CountryName)
	if err != nil {
		panic("Error in parse Channel")
	}
	return c, nil
}

func CreateHumanDatCollectionRes(rows *sql.Rows) ([]HumanDat, error) {
	var humanDatCollection []HumanDat
	for rows.Next() {
		j, err := parseHumanDatResponse(rows)
		if err != nil {
			panic("Error in parsing array of Channels")
		}
		humanDatCollection = append(humanDatCollection, j)
	}
	return humanDatCollection, nil
}
