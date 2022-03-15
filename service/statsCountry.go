package service
//
//import (
//	"github.com/Pivnoy/isbd/models"
//	"github.com/Pivnoy/isbd/server"
//)
//
//func getStatsCountry(tableName string) (models.StatsCountryResponse, error) {
//	rows, err := server.DbInstance.Query("select gep from economics where country_name = $1", tableName)
//	if err != nil {
//		panic("Error in going to crazy table")
//	}
//
//	rows, err = server.DbInstance.Query("select salary from human where country_name = $1", tableName)
//
//	return models.StatsCountryResponse{}, err
//}
