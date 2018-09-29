package repository

import (
	"server/model"
)

func GetUserByEmail(email string) *model.User {
	db := openDBConnection()
	defer db.Close()

	stmt, err := db.Prepare("SELECT id, name, email, password FROM users WHERE deleted_at IS NULL AND email = (?)")
	defer stmt.Close()
	check(err)

	var row, e = stmt.Query(email)
	defer row.Close()
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

func CreateUser(name string, email string, password string) (int64, error) {
	db := openDBConnection()
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO users(name, email, password, created_at, updated_at) VALUES((?), (?), (?), CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)")
	defer stmt.Close()
	check(err)

	result, err := stmt.Exec(name, email, password)
	var resultedID int64
	if err == nil {
		resultedID, err = result.LastInsertId()
	}
	return resultedID, err
}

func GetUserByID(id int64) *model.User {
	db := openDBConnection()
	defer db.Close()

	stmt, err := db.Prepare("SELECT id, name, email, password FROM users WHERE deleted_at IS NULL AND id = (?)")
	defer stmt.Close()
	check(err)

	var row, e = stmt.Query(id)
	defer row.Close()
	check(e)
	if row.Next() {
		var id int64
		var name, email, password *string
		err = row.Scan(&id, &name, &email, &password)
		check(err)

		return &model.User{id, name, email, password}
	}

	return nil
}
