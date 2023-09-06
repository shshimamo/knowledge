package service

import (
	"context"
	"errors"
	gql "github.com/shshimamo/knowledge/main/graph/model"
	"github.com/shshimamo/knowledge/main/middlewares"
	"github.com/shshimamo/knowledge/main/model"
	"github.com/shshimamo/knowledge/main/repository"
)

type UserService interface {
	CreateUser(ctx context.Context, input *gql.NewUser) (*gql.User, error)
	GetUser(ctx context.Context, id int) (*gql.User, error)
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

func (s *userService) GetUser(ctx context.Context, id int) (*gql.User, error) {
	u, err := s.userRepo.GetUser(ctx, &repository.GetUserCommand{ID: id})
	if err != nil {
		return nil, err
	}

	return model.MapUserModelToGql(u), nil
}
