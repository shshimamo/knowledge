package model

type CurrentUser struct {
	ID         string       `json:"id"`
	AuthUserID string       `json:"authUserId"`
	Name       *string      `json:"name,omitempty"`
	Knowledge  *Knowledge   `json:"knowledge"`
	Knowledges []*Knowledge `json:"knowledges"`
}
