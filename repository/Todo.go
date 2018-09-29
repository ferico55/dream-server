package repository

import (
	"server/model"
	"strconv"
)

func GetTodoOwnerID(id string) int64 {
	db := openDBConnection()
	defer db.Close()

	stmt, err := db.Prepare("SELECT d.user_id FROM todos t JOIN dreams d ON t.dream_id = d.id WHERE t.id = (?) AND d.deleted_at IS NULL AND t.deleted_at IS NULL")
	defer stmt.Close()
	check(err)

	row, err := stmt.Query(id)
	check(err)
	defer row.Close()

	var userID *string
	row.Next()
	err = row.Scan(&userID)
	check(err)

	if userID == nil {
		return 0
	}
	userIDInt, err := strconv.ParseInt(*userID, 10, 64)
	check(err)

	return userIDInt
}

func CheckTodo(id string) error {
	db := openDBConnection()
	defer db.Close()

	stmt, err := db.Prepare("UPDATE todos SET is_checked = '1', updated_at = CURRENT_TIMESTAMP WHERE id = (?)")
	defer stmt.Close()
	check(err)

	_, err = stmt.Exec(id)
	return err
}

func UncheckTodo(id string) error {
	db := openDBConnection()
	defer db.Close()

	stmt, err := db.Prepare("UPDATE todos SET is_checked = '0', updated_at = CURRENT_TIMESTAMP WHERE id = (?)")
	defer stmt.Close()
	check(err)

	_, err = stmt.Exec(id)
	return err
}

func CreateTodo(title string, dreamID string) (int64, error) {
	db := openDBConnection()
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO todos(dream_id, title, is_checked, created_at, updated_at) VALUES((?), (?), 0, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)")
	defer stmt.Close()
	check(err)

	result, err := stmt.Exec(dreamID, title)
	var resultedID int64
	if err == nil {
		resultedID, err = result.LastInsertId()
	}
	return resultedID, err
}

func GetTodoByID(id int64) *model.Todo {
	db := openDBConnection()
	defer db.Close()

	stmt, err := db.Prepare("SELECT id, title, is_checked FROM todos WHERE id = (?) AND deleted_at IS NULL")
	defer stmt.Close()
	check(err)

	row, err := stmt.Query(id)
	check(err)
	defer row.Close()

	var todoID, is_checked int
	var title string
	row.Next()
	err = row.Scan(&todoID, &title, &is_checked)
	check(err)

	return &model.Todo{todoID, title, is_checked == 0}
}
