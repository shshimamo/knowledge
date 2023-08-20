package service

import (
	"context"
	"database/sql"
	"errors"
	"github.com/shshimamo/knowledge-main/middlewares/auth"
	"github.com/shshimamo/knowledge-main/model"
	"github.com/shshimamo/knowledge-main/repository"

	gql "github.com/shshimamo/knowledge-main/graph/model"
)

type UserService interface {
	CreateUser(ctx context.Context, input *gql.NewUser) (*gql.User, error)
}

type userService struct {
	db *sql.DB
}

func newUserService(db *sql.DB) *userService {
	return &userService{db}
}

func (s *userService) CreateUser(ctx context.Context, gqlnew *gql.NewUser) (*gql.User, error) {
	token, ok := model.GetCurrentToken(ctx)
	if !ok {
		return nil, errors.New("not authenticated")
	}
	_, ok = auth.GetCurrentUser(ctx)
	if ok {
		return nil, errors.New("Already registered")
	}

	user := model.MapUserGqlNewToModel(gqlnew)
	user.AuthUserID = token.AuthUserID
	err := user.Validate()
	if err != nil {
		return nil, err
	}

	repo := repository.NewUserRepository(s.db)
	newuser, err := repo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	gqluser := model.MapUserModelToGql(newuser)

	return gqluser, nil
}
