package service

import (
	"github.com/Pivnoy/isbd/models"
	"github.com/Pivnoy/isbd/server"
)

func GetSects() (models.SectsResponse, error) {
	rows, err := server.DbInstance.Query("select sect.id, sect.name from sect")
	if err != nil {
		panic("Error in going to sect table")
	}
	res, err := models.CreateSectsCollection(rows)
	if err != nil {
		return models.SectsResponse{}, err
	}
	return models.SectsResponse{Sects: res}, nil
}
