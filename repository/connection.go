package repository

import (
	"database/sql"
	"fmt"
)

var db *sql.DB

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
