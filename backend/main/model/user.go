package model

import (
	"github.com/shshimamo/knowledge-main/db"
	gqlModel "github.com/shshimamo/knowledge-main/graph/model"
	"github.com/volatiletech/null/v8"
	"strconv"
)

type User struct {
	ID         int
	AuthUserID int
	Name       string
}

func MapUserDBToModel(dbuser *db.User) *User {
	user := &User{
		ID: dbuser.ID,
	}
	if dbuser.AuthUserID.Valid {
		user.AuthUserID = int(dbuser.AuthUserID.Int64)
	}
	if dbuser.Name.Valid {
		user.Name = dbuser.Name.String
	}
	return user
}

func MapUserModelToDB(user *User) *db.User {
	dbuser := &db.User{
		AuthUserID: null.Int64From(int64(user.AuthUserID)),
	}
	if user.Name != "" {
		dbuser.Name = null.StringFrom(user.Name)
	}
	return dbuser
}

func MapNewUserGraphToModel(newuser *gqlModel.NewUser) *User {
	user := &User{}
	user.Name = newuser.Name
	return user
}

func MapUserModelToGraph(user *User) *gqlModel.User {
	var name *string
	if user.Name != "" {
		temp := user.Name
		name = &temp
	}
	gqluser := &gqlModel.User{
		ID:         strconv.Itoa(user.ID),
		AuthUserID: strconv.Itoa(user.AuthUserID),
		Name:       name,
	}
	return gqluser
}