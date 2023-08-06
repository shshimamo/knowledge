package model

type Knowledge struct {
	ID          string `json:"id"`
	UserID      string `json:"userId"`
	Title       string `json:"title"`
	Text        string `json:"text"`
	IsPublic    bool   `json:"isPublic"`
	PublishedAt string `json:"publishedAt"`
}
