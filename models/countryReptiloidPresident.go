package models

import "database/sql"

type CountryReptiloidName struct {
	CountryName string
}

type CountryReptiloidId struct {
	Id int
}

func parseCountryReptiloidName(rows *sql.Rows) (CountryReptiloidName, error) {
	c := CountryReptiloidName{}
	err := rows.Scan(&c.CountryName)
	if err != nil {
		panic("Error in parse Channel")
	}
	return c, nil
}

func CreateCountryReptiloidNameCollection(rows *sql.Rows) ([]CountryReptiloidName, error) {
	var countryReptiloidCollection []CountryReptiloidName
	for rows.Next() {
		j, err := parseCountryReptiloidName(rows)
		if err != nil {
			panic("Error in parsing array of tv work")
		}
		countryReptiloidCollection = append(countryReptiloidCollection, j)
	}
	return countryReptiloidCollection, nil
}

func parseCountryReptiloidId(rows *sql.Rows) (CountryReptiloidId, error) {
	c := CountryReptiloidId{}
	err := rows.Scan(&c.Id)
	if err != nil {
		panic("Error in parse Channel")
	}
	return c, nil
}

func CreateCountryReptiloidIdCollection(rows *sql.Rows) ([]CountryReptiloidId, error) {
	var countryReptiloidIdCollection []CountryReptiloidId
	for rows.Next() {
		j, err := parseCountryReptiloidId(rows)
		if err != nil {
			panic("Error in parsing array of tv work")
		}
		countryReptiloidIdCollection = append(countryReptiloidIdCollection, j)
	}
	return countryReptiloidIdCollection, nil
}
