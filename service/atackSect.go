package service

import "github.com/Pivnoy/isbd/server"

func DoAttackSect(sectName, punitiveName string) error {
	_, err := server.DbInstance.Query("update punitive_structure set sect_name = $1 where name = $2", sectName, punitiveName)
	if err != nil {
		panic("Error to going punitive structure table")
	}
	return nil
}
