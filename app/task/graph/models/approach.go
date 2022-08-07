package models

type Approach struct {
	ID        string `json:"id"`
	Content   string `json:"content"`
	VoteCount int    `json:"voteCount"`
	CreatedAt string `json:"createdAt"`
}

func (Approach) IsSearchResultItem() {}
