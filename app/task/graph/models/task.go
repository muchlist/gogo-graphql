package models

type Task struct {
	ID            string   `json:"id"`
	Content       string   `json:"content"`
	Tags          []string `json:"tags"`
	ApproachCount int      `json:"approachCount"`
	CreatedAt     string   `json:"createdAt"`
}

func (Task) IsSearchResultItem() {}
