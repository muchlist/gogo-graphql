package models

type Approach struct {
	ID        int64  `json:"id"`
	Content   string `json:"content"`
	VoteCount int    `json:"voteCount"`
	CreatedAt string `json:"createdAt"`
}

func (Approach) IsSearchResultItem() {}
