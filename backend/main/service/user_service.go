package service

import (
	"context"
	"database/sql"
	"errors"
	"github.com/shshimamo/knowledge-main/model"
	"github.com/shshimamo/knowledge-main/repository"

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

func (s *userService) CreateUser(ctx context.Context, input *graphModel.NewUser) (*graphModel.User, error) {
	token, ok := model.GetCurrentToken(ctx)
	if !ok {
		return nil, errors.New("not authenticated")
	}
	_, ok = model.GetCurrentUser(ctx)
	if ok {
		return nil, errors.New("Already registered")
	}

	user := model.MapNewUserGraphToModel(input)
	user.AuthUserID = token.AuthUserID

	repo := repository.NewUserRepository(s.db)
	newuser, err := repo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	gqluser := model.MapUserModelToGraph(newuser)

	return gqluser, nil
}
