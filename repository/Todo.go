package repository

import "strconv"

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
