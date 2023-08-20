package model

import (
	"github.com/shshimamo/knowledge-main/db"
	gql "github.com/shshimamo/knowledge-main/graph/model"
	"github.com/shshimamo/knowledge-main/model/errs"
	"github.com/volatiletech/null/v8"
	"strconv"
)

type User struct {
	ID         int
	AuthUserID int
	Name       string
}

func (u *User) Validate() error {
	if u.AuthUserID == 0 {
		return errs.NewValidationError("AuthUserID is required")
	}
	if u.Name == "" {
		return errs.NewValidationError("Name is required")
	}
	return nil
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

func MapUserGqlNewToModel(gqlnew *gql.NewUser) *User {
	u := &User{}
	u.Name = gqlnew.Name
	return u
}

func MapUserModelToDB(u *User) *db.User {
	db := &db.User{
		AuthUserID: null.Int64From(int64(u.AuthUserID)),
	}
	if u.Name != "" {
		db.Name = null.StringFrom(u.Name)
	}
	return db
}

func MapUserModelToGql(u *User) *gql.User {
	var name *string
	if u.Name != "" {
		temp := u.Name
		name = &temp
	}
	gql := &gql.User{
		ID:         strconv.Itoa(u.ID),
		AuthUserID: strconv.Itoa(u.AuthUserID),
		Name:       name,
	}
	return gql
}

func MapUserModelToGqlCurrent(u *User) *gql.CurrentUser {
	var name *string
	if u.Name != "" {
		temp := u.Name
		name = &temp
	}
	gql := &gql.CurrentUser{
		ID:         strconv.Itoa(u.ID),
		AuthUserID: strconv.Itoa(u.AuthUserID),
		Name:       name,
	}
	return gql
}
