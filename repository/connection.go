package repository

import (
	"database/sql"
	"fmt"
	"server/config"
)

func openDBConnection() *sql.DB {
	db, err := sql.Open(config.DriverName, config.ConnectionString)
	check(err)
	return db
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
