package models

import "database/sql"

type CountryResponse struct {
	Countrys []Country `json:"countrys"`
}

type Country struct {
	Id   int
	Name string
}

func parseCountry(rows *sql.Rows) (Country, error) {
	c := Country{}
	err := rows.Scan(&c.Id, &c.Name)
	if err != nil {
		panic("Error in parse Channel")
	}
	return c, nil
}

func CreateCountryCollection(rows *sql.Rows) ([]Country, error) {
	var countryCollection []Country
	for rows.Next() {
		j, err := parseCountry(rows)
		if err != nil {
			panic("Error in parsing array of Channels")
		}
		countryCollection = append(countryCollection, j)
	}
	return countryCollection, nil
}
