package service

import (
	"github.com/Pivnoy/isbd/models"
	"github.com/Pivnoy/isbd/server"
)

func GetPunitive() (models.PunitiveResponse, error) {
	rows, err := server.DbInstance.Query("select punitive_structure.id, punitive_structure.name from punitive_structure")
	if err != nil {
		panic("Error in going to punitive_structure table")
	}
	res, err := models.CreatePunitiveCollection(rows)
	if err != nil {
		return models.PunitiveResponse{}, err
	}
	return models.PunitiveResponse{Structures: res}, nil
}
