package model

type User struct {
	ID         string  `json:"id"`
	AuthUserID string  `json:"authUserId"`
	Name       *string `json:"name,omitempty"`
}
