package service

import (
	"context"
	"errors"
	"github.com/shshimamo/knowledge-main/middlewares"
	"github.com/shshimamo/knowledge-main/model"
	"github.com/shshimamo/knowledge-main/repository"

	gql "github.com/shshimamo/knowledge-main/graph/model"
)

type UserService interface {
	CreateUser(ctx context.Context, input *gql.NewUser) (*gql.User, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func newUserService(userRepo repository.UserRepository) *userService {
	return &userService{userRepo: userRepo}
}

func (s *userService) CreateUser(ctx context.Context, gqlNew *gql.NewUser) (*gql.User, error) {
	token, ok := middlewares.GetCurrentToken(ctx)
	if !ok {
		return nil, errors.New("not authenticated")
	}
	_, ok = middlewares.GetCurrentUser(ctx)
	if ok {
		return nil, errors.New("Already registered")
	}

	user := model.MapUserGqlNewToModel(gqlNew)
	user.AuthUserID = token.AuthUserID
	err := user.Validate()
	if err != nil {
		return nil, err
	}

	newUser, err := s.userRepo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	gqlUser := model.MapUserModelToGql(newUser)

	return gqlUser, nil
}
