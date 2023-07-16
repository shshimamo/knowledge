package model

type AuthUser struct {
	ID             int    `json:"-"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	PasswordDigest string `json:"-"`
}
