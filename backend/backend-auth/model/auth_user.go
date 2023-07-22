package model

import "golang.org/x/crypto/bcrypt"

type AuthUser struct {
	ID             int
	Email          string
	PasswordDigest string
}

func NewAuthUser(email string) *AuthUser {
	return &AuthUser{
		Email: email,
	}
}

func (au *AuthUser) CreatePasswordDigest(password string) error {
	digest, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	au.PasswordDigest = string(digest)
	return nil
}
