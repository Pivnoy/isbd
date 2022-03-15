package service

import "github.com/Pivnoy/isbd/server"

func DoMorth(human_id, reptiloid_id int) error {
	_, err := server.DbInstance.Query("update reptiloid set human_id = $1 where id = $2", human_id, reptiloid_id)
	if err != nil {
		panic("Error in going to crazy table")
	}
	return nil
}
