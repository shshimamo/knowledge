package model

import (
	"context"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

type Token struct {
	AuthUserID int
}

type claims struct {
	AuthUserID int `json:"auth_user_id"`
	jwt.StandardClaims
}

func NewToken(tokenStr string) (*Token, error) {
	token, err := getJWTToken(tokenStr)
	if err != nil {
		return nil, err
	}

	claims, err := getClaims(token)
	if err != nil {
		return nil, err
	}

	return &Token{
		AuthUserID: claims.AuthUserID,
	}, nil
}

func getJWTToken(tokenStr string) (*jwt.Token, error) {
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

	return token, err
}

func getClaims(token *jwt.Token) (*claims, error) {
	claims, ok := token.Claims.(*claims)
	if ok && token.Valid && claims.AuthUserID != 0 {
		return claims, nil
	} else {
		return nil, errors.New("Invalid token")
	}
}

type CurrentTokenKey struct{}

func GetCurrentToken(ctx context.Context) (*Token, bool) {
	switch v := ctx.Value(CurrentTokenKey{}).(type) {
	case *Token:
		return v, true
	default:
		return nil, false
	}
}
