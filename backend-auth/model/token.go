package model

type Token struct {
	ID         int
	AuthUserID int
	Token      string
	Active     bool
}
