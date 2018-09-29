package repository

import (
	"database/sql"
	"server/config"
	"server/model"
)

func GetUserByEmail(email string) *model.User {
	db, err := sql.Open(config.DriverName, config.ConnectionString)
	check(err)
	defer db.Close()

	stmt, err := db.Prepare("SELECT id, name, email, password FROM users WHERE deleted_at IS NULL AND email = (?)")
	defer stmt.Close()
	check(err)

	var row, e = stmt.Query(email)
	check(e)
	if row.Next() {
		var id int64
		var name, email, password *string
		err = row.Scan(&id, &name, &email, &password)
		check(err)

		var user = model.User{id, name, email, password}
		return &user
	}

	return nil
}
