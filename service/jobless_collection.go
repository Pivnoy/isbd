package service

import (
	"github.com/Pivnoy/isbd/models"
	"github.com/Pivnoy/isbd/server"
)

func GetAllJobless() (models.JoblessResponse, error) {
	rows, err := server.DbInstance.Query("select * from human")
	if err != nil {
		panic("Error in going to jobless table")
	}
	res, err := models.CreateJoblessCollection(rows)
	if err != nil {
		return models.JoblessResponse{}, err
	}
	var count int // кол-во безработных
	for _, jobless := range res {
		if jobless.Employment == false {
			count += 1
		}
	}
	ver := float32(count) / float32(len(res)) // процентное соотношение
	return models.JoblessResponse{Jobless: ver, Jobless_c: count}, nil
}
