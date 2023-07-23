package model

import "strconv"

type Token struct {
	AuthUserID int
}

func NewToken(authUserID string) (*Token, error) {
	ID, err := strconv.Atoi(authUserID)
	if err != nil {
		return nil, err
	}

	return &Token{
		AuthUserID: ID,
	}, nil
}
