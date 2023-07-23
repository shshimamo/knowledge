package model

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Token struct {
	ID         int
	AuthUserID int
	Token      TokenStr
	Active     bool
	ExpiresAt  time.Time
}

type TokenStr string

type claims struct {
	AuthUserID int `json:"auth_user_id"`
	jwt.StandardClaims
}

func NewToken(auser *AuthUser) *Token {
	expiresAt := time.Now().Add(time.Hour * 24 * 7)
	//token := jwt.New(jwt.SigningMethodHS256)
	claims := &claims{
		AuthUserID: auser.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, _ := token.SignedString([]byte("secret")) // TODO: 環境変数にする

	return &Token{
		AuthUserID: auser.ID,
		Token:      TokenStr(tokenStr),
		Active:     true,
		ExpiresAt:  expiresAt,
	}
}
