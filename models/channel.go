package models

import (
	"database/sql"
)

type Channel struct {
	Id     uint64
	Number int
	Name   string
	Theme  string
}

func parseChannel(rows *sql.Rows) (Channel, error) {
	c := Channel{}
	err := rows.Scan(&c.Id, &c.Number, &c.Name, &c.Theme)
	if err != nil {
		panic("Error in parse Channel")
	}
	return c, nil
}

func CreateChannelCollection(rows *sql.Rows) []Channel {
	var channelCollection []Channel
	for rows.Next() {
		c, err := parseChannel(rows)
		if err != nil {
			panic("Error in parsing array of Channels")
		}
		channelCollection = append(channelCollection, c)
	}
	return channelCollection
}
