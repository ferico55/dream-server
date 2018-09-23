package repository

import (
	"database/sql"
	"fmt"
	"server/config"
)

var db *sql.DB

func openDBConnection() *sql.DB {
	var err error

	db, err = sql.Open(config.DriverName, config.ConnectionString)
	check(err)

	return db
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
