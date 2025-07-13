package usecase

import (
	"context"
	"errors"
	gql "github.com/shshimamo/knowledge/main/graph/model"
	"github.com/shshimamo/knowledge/main/middlewares"
	"github.com/shshimamo/knowledge/main/model"
	"github.com/shshimamo/knowledge/main/repository"
)

type UserUseCase interface {
	CreateUser(ctx context.Context, input *gql.NewUser) (*gql.User, error)
	GetUser(ctx context.Context, id int64) (*gql.User, error)
}

type userUseCase struct {
	userRepo repository.UserRepository
}

func NewUserUseCase(userRepo repository.UserRepository) UserUseCase {
	return &userUseCase{userRepo: userRepo}
}

func (u *userUseCase) CreateUser(ctx context.Context, gqlNew *gql.NewUser) (*gql.User, error) {
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

	newUser, err := u.userRepo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	gqlUser := model.MapUserModelToGql(newUser)

	return gqlUser, nil
}

func (u *userUseCase) GetUser(ctx context.Context, id int64) (*gql.User, error) {
	user, err := u.userRepo.GetUser(ctx, &repository.GetUserCommand{ID: id})
	if err != nil {
		return nil, err
	}

	return model.MapUserModelToGql(user), nil
}