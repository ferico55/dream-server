package repository

import (
	"fmt"
	"server/model"
)

func GetAllDreams() []model.Dream {
	db := openDBConnection()
	defer db.Close()

	row, err := db.Query("SELECT id, user_id, title, description, image_uri FROM dreams WHERE deleted_at IS NULL")
	defer row.Close()
	check(err)

	var dreams []model.Dream
	var dream *model.Dream
	var prevID int = 0
	for row.Next() {
		var id, userID, todoID, isChecked int
		var title, description, imageURI, todoTitle *string
		err = row.Scan(&id, &userID, &title, &description, &imageURI, &todoID, &todoTitle, &isChecked)
		check(err)

		if id == prevID {
			todo := model.Todo{todoID, *todoTitle, (isChecked == 1)}
			dream.Todo = append(dream.Todo, todo)
		} else {
			if dream != nil {
				dreams = append(dreams, *dream)
			}
			todos := make([]model.Todo, 0)
			dream = &model.Dream{id, userID, title, description, imageURI, todos}

			if todoTitle != nil {
				todo := model.Todo{todoID, *todoTitle, (isChecked == 1)}
				dream.Todo = append(dream.Todo, todo)
			}
			prevID = id
		}
	}

	if dream != nil {
		dreams = append(dreams, *dream)
	}

	return dreams
}

func GetMyDreams(uid int) []model.Dream {
	db := openDBConnection()
	defer db.Close()

	stmt, err := db.Prepare("SELECT d.id, d.user_id, d.title, d.description, d.image_uri, t.id AS todo_id, t.title as todo_title, t.is_checked FROM dreams d LEFT JOIN todos t ON d.id = t.dream_id WHERE d.deleted_at IS NULL AND t.deleted_at IS NULL AND user_id = (?)")
	defer stmt.Close()
	check(err)

	row, err := stmt.Query(uid)
	check(err)
	defer row.Close()

	var dreams []model.Dream
	var dream *model.Dream
	var prevID int
	for row.Next() {
		var id, userID, todoID, isChecked int
		var title, description, imageURI, todoTitle *string
		err = row.Scan(&id, &userID, &title, &description, &imageURI, &todoID, &todoTitle, &isChecked)
		check(err)

		if id == prevID {
			todo := model.Todo{todoID, *todoTitle, (isChecked == 1)}
			fmt.Println("asdf")
			fmt.Println(todo)
			dream.Todo = append(dream.Todo, todo)
		} else {
			if dream != nil {
				dreams = append(dreams, *dream)
			}
			todos := make([]model.Todo, 0)
			dream = &model.Dream{id, userID, title, description, imageURI, todos}

			if todoTitle != nil {
				todo := model.Todo{todoID, *todoTitle, (isChecked == 1)}
				dream.Todo = append(dream.Todo, todo)
			}
			prevID = id
		}
	}

	if dream != nil {
		dreams = append(dreams, *dream)
	}

	return dreams
}

func GetDreamByID(id string) *model.Dream {
	db := openDBConnection()
	defer db.Close()

	stmt, err := db.Prepare("SELECT d.id, d.user_id, d.title, d.description, d.image_uri, t.id AS todo_id, t.title as todo_title, t.is_checked FROM dreams d LEFT JOIN todos t ON d.id = t.dream_id WHERE d.deleted_at IS NULL AND t.deleted_at IS NULL AND d.id = (?)")
	defer stmt.Close()
	check(err)

	row, err := stmt.Query(id)
	defer row.Close()
	check(err)

	var dream *model.Dream
	var prevID int
	for row.Next() {
		var id, userID, todoID, isChecked int
		var title, description, imageURI, todoTitle *string
		err = row.Scan(&id, &userID, &title, &description, &imageURI, &todoID, &todoTitle, &isChecked)
		check(err)

		if id == prevID {
			todo := model.Todo{todoID, *todoTitle, (isChecked == 1)}
			dream.Todo = append(dream.Todo, todo)
		} else {
			todos := make([]model.Todo, 0)
			dream = &model.Dream{id, userID, title, description, imageURI, todos}

			if todoTitle != nil {
				todo := model.Todo{todoID, *todoTitle, (isChecked == 1)}
				dream.Todo = append(dream.Todo, todo)
			}
			prevID = id
		}
	}

	return dream
}

func CreateDream(title string, description string, imageURI string, userID int64) (int64, error) {
	db := openDBConnection()
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO dreams(user_id, title, description, image_uri, created_at, updated_at) VALUES((?), (?), (?), (?), CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)")
	defer stmt.Close()
	check(err)

	result, err := stmt.Exec(userID, title, description, imageURI)
	var resultedID int64
	if err == nil {
		resultedID, err = result.LastInsertId()
	}
	return resultedID, err
}

func GetDreamOwnerID(id string) int {
	db := openDBConnection()
	defer db.Close()

	stmt, err := db.Prepare("SELECT user_id FROM dreams WHERE id = (?) AND deleted_at IS NULL")
	defer stmt.Close()
	check(err)

	row, err := stmt.Query(id)
	defer row.Close()
	check(err)

	var userID int
	row.Next()
	err = row.Scan(&userID)
	check(err)

	return userID
}
