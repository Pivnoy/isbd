package service

import (
	"github.com/Pivnoy/isbd/models"
	"github.com/Pivnoy/isbd/server"
)

func GetCountry() (models.CountryResponse, error) {
	rows, err := server.DbInstance.Query("select country.id, country.name from country")
	if err != nil {
		panic("Error in going to punitive_structure table")
	}
	res, err := models.CreateCountryCollection(rows)
	if err != nil {
		return models.CountryResponse{}, err
	}
	return models.CountryResponse{Countrys: res}, nil
}
