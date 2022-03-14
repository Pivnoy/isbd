package service

import (
	"github.com/Pivnoy/isbd/models"
	"github.com/Pivnoy/isbd/server"
)

func GetAllJobless() ([]models.Jobless, error) {
	rows, err := server.DbInstance.Query("select * from jobless")
	if err != nil {
		panic("Error in going to jobless table")
	}
	res, err := models.CreateJoblessCollection(rows)
	if err != nil {
		return []models.Jobless{}, err
	}
	return res, nil
}
