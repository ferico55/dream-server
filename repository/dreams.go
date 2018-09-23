package repository

import (
	"server/model"
)

func GetAllDreams() []model.Dream {
	var db = openDBConnection()
	defer db.Close()

	var rows, err = db.Query("SELECT id, user_id, title, description, image_uri FROM dreams WHERE deleted_at IS NULL")
	check(err)

	var dreams []model.Dream
	var id, userID int
	var title, description, imageURI *string
	for rows.Next() {
		err = rows.Scan(&id, &userID, &title, &description, &imageURI)
		check(err)
		var dream = model.Dream{id, userID, title, description, imageURI}
		dreams = append(dreams, dream)
	}

	return dreams
}
