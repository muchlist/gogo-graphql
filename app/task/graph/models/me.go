package models

type Me struct {
	ID       string  `json:"id"`
	Username *string `json:"username"`
	Name     string  `json:"name"`
	TaskList []*Task `json:"taskList"`
}
