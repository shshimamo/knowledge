package model

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

type Token struct {
	AuthUserID int
}

func NewToken(tokenStr string) (*Token, error) {
	claims, err := getClaims(tokenStr)
	if err != nil {
		return nil, err
	}

	return &Token{
		AuthUserID: claims.AuthUserID,
	}, nil
}

type claims struct {
	AuthUserID int `json:"auth_user_id"`
	jwt.StandardClaims
}

func getClaims(tokenStr string) (*claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &claims{}, func(token *jwt.Token) (interface{}, error) {
		// Always make sure the token method corresponds to the one you expect.
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// Return the key for validation
		return []byte("secret"), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*claims)
	if ok && token.Valid && claims.AuthUserID != 0 {
		return claims, nil
	} else {
		return nil, errors.New("Invalid token")
	}
}
