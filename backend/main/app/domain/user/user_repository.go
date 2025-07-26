package user

import (
	"context"
)

type UserRepository interface {
	GetUserByID(ctx context.Context, userID int64) (*User, error)
	GetUserByUserName(ctx context.Context, userName string) (*User, error)
	CreateUser(ctx context.Context, u *User) (*User, error)
	UpdateUser(ctx context.Context, u *User) (*User, error)
	DeleteUser(ctx context.Context, u *User) error
}