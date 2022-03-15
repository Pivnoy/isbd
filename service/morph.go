package service

import "github.com/Pivnoy/isbd/server"

func DoMorth(humanId, reptiloidId int) error {
	_, err := server.DbInstance.Query("update reptiloid set human_id = $1 where id = $2", humanId, reptiloidId)
	if err != nil {
		panic("Error in going to crazy table")
	}
	return nil
}
