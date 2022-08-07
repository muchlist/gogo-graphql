package models

type Task struct {
	ID            int64    `json:"id"`
	Content       string   `json:"content"`
	Tags          []string `json:"tags"`
	ApproachCount int      `json:"approachCount"`
	CreatedAt     string   `json:"createdAt"`
}

func (Task) IsSearchResultItem() {}
