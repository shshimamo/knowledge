package model

import (
	"github.com/shshimamo/knowledge-main/db"
)

type User struct {
	ID         int
	AuthUserId int
	Name       string
}

func ConvertUserFromDB(dbUser *db.User) *User {
	newUser := &User{
		ID: dbUser.ID,
	}
	if dbUser.AuthUserID.Valid {
		newUser.AuthUserId = int(dbUser.AuthUserID.Int64)
	}
	if dbUser.Name.Valid {
		newUser.Name = dbUser.Name.String
	}
	return newUser
}
