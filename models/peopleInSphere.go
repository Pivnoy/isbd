package models

import "database/sql"

type PeopleInSphere struct {
	HumanId int
}

func parsePeopleInSphere(rows *sql.Rows) (PeopleInSphere, error) {
	c := PeopleInSphere{}
	err := rows.Scan(&c.HumanId)
	if err != nil {
		panic("Error in parse Channel")
	}
	return c, nil
}

func CreatePeopleInSphereCollectionCollection(rows *sql.Rows) ([]PeopleInSphere, error) {
	var peopleInSphereCollection []PeopleInSphere
	for rows.Next() {
		j, err := parsePeopleInSphere(rows)
		if err != nil {
			panic("Error in parsing array of tv work")
		}
		peopleInSphereCollection = append(peopleInSphereCollection, j)
	}
	return peopleInSphereCollection, nil
}
