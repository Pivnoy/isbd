package service

import (
	"github.com/Pivnoy/isbd/models"
	"github.com/Pivnoy/isbd/server"
)

func GetCrazyAll() (models.CrazyResponse, error) {
	rows, err := server.DbInstance.Query("select crazy.id, crazy.human_id from crazy")
	if err != nil {
		panic("Error in going to crazy table")
	}
	res, err := models.CreateCrazyCollection(rows)
	if err != nil {
		return models.CrazyResponse{}, err
	}

	rows, err = server.DbInstance.Query("select * from human")
	if err != nil {
		panic("Error in going human table")
	}
	hum, err := models.CreateHumanCollection(rows)

	if err != nil {
		return models.CrazyResponse{}, err
	}

	return models.CrazyResponse{Crazy: float32(len(res)) / float32(len(hum)), Crazy_c: len(res)}, nil
}
