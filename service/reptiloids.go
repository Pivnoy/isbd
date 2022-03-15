package service

import (
	"github.com/Pivnoy/isbd/models"
	"github.com/Pivnoy/isbd/server"
)

func GetReptioids() (models.ReptioloidResponse, error) {
	rows, err := server.DbInstance.Query("select id, name from reptiloid")
	if err != nil {
		panic("Error in going to punitive_structure table")
	}
	res, err := models.CreateReptiloidCollection(rows)
	if err != nil {
		return models.ReptioloidResponse{}, err
	}
	return models.ReptioloidResponse{Reptiloids: res}, nil
}
