package service

import (
	"github.com/Pivnoy/isbd/models"
	"github.com/Pivnoy/isbd/server"
)

func GetHuman() (models.HumanDateResponse, error) {
	rows, err := server.DbInstance.Query("select id, name, country_name from human ")
	if err != nil {
		panic("Error in going to punitive_structure table")
	}
	res, err := models.CreateHumanDatCollectionRes(rows)
	if err != nil {
		return models.HumanDateResponse{}, err
	}
	return models.HumanDateResponse{Humans: res}, nil

}
