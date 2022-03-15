package models

import "database/sql"

type ReptiloidsResponseAll struct {
	Tv         float32                `json:"tv"`
	Business   float32                `json:"business"`
	Science    float32                `json:"science"`
	Countries  []CountryReptiloidName `json:"countries"`
	CountriesP float32                `json:"countries_p"`
}

type RepriloidResponseCountry struct {
	Tv       float32 `json:"tv"`
	Business float32 `json:"business"`
	Science  float32 `json:"science"`
}

type ReptioloidResponse struct {
	Reptiloids []Reptiloid `json:"reptiloids"`
}

type Reptiloid struct {
	Id   int
	Name string
}

func parseReptiloid(rows *sql.Rows) (Reptiloid, error) {
	c := Reptiloid{}
	err := rows.Scan(&c.Id, &c.Name)
	if err != nil {
		panic("Error in parse Channel")
	}
	return c, nil
}

func CreateReptiloidCollection(rows *sql.Rows) ([]Reptiloid, error) {
	var reptiloidCollection []Reptiloid
	for rows.Next() {
		j, err := parseReptiloid(rows)
		if err != nil {
			panic("Error in parsing array of Channels")
		}
		reptiloidCollection = append(reptiloidCollection, j)
	}
	return reptiloidCollection, nil
}
