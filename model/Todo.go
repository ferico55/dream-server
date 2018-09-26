package model

type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	IsChecked bool   `json:"is_checked"`
}
