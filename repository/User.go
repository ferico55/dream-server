package repository

import "server/model"

func GetUserByEmailAndPassword(email string, password string) *model.User {
	var db = openDBConnection()
	defer db.Close()

	// var rows, err = db.Query("SELECT id, name, email FROM users WHERE deleted_at IS NULL")
	// check(err)
	return nil
}
