package service

// добавить количество всех людей
import (
	"github.com/Pivnoy/isbd/models"
	"github.com/Pivnoy/isbd/server"
)

func GetJoblessCountry(countryName string) (models.JoblessResponse, error) {
	rows, err := server.DbInstance.Query("select * from human where country_name = $1", countryName)
	if err != nil {
		panic("Error in going to human table")
	}
	res, err := models.CreateHumanCollection(rows)
	if err != nil {
		return models.JoblessResponse{}, err
	}
	var count int // кол-во безработных
	for _, jobless := range res {
		if jobless.Employment == false {
			count += 1
		}
	}

	return models.JoblessResponse{Jobless: float32(count) / float32(len(res)), JoblessC: count, PeopleAmount: len(res)}, nil
}
