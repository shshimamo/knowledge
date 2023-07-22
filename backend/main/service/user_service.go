package service

import (
	"context"
	"database/sql"
	graphModel "github.com/shshimamo/knowledge-main/graph/model"
)

type UserService interface {
	CreateUser(ctx context.Context, input *graphModel.NewUser) (*graphModel.User, error)
}

type userService struct {
	db *sql.DB
}

func NewUserService(db *sql.DB) *userService {
	return &userService{db}
}

func (u *userService) CreateUser(ctx context.Context, input *graphModel.NewUser) (*graphModel.User, error) {
	// TODO: implement
	return nil, nil
}
