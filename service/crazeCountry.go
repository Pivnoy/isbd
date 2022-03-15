package service

import (
	"github.com/Pivnoy/isbd/models"
	"github.com/Pivnoy/isbd/server"
)

func GetCrazyCountry(country_name string) (models.CrazyResponse, error) {
	rows, err := server.DbInstance.Query("select crazy.id, crazy.human_id from crazy inner join human h on h.id = crazy.human_id where h.country_name = $1", country_name)
	if err != nil {
		panic("Error in going to crazy table")
	}
	res, err := models.CreateCrazyCollection(rows)
	if err != nil {
		return models.CrazyResponse{}, err
	}

	rows, err = server.DbInstance.Query("select * from human where country_name = $1", country_name)
	if err != nil {
		panic("Error in going human table")
	}
	hum, err := models.CreateHumanCollection(rows)

	if err != nil {
		return models.CrazyResponse{}, err
	}

	return models.CrazyResponse{Crazy: float32(len(res)) / float32(len(hum)), Crazy_c: len(res)}, nil
}
