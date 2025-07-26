package user

import (
	"context"
	"errors"

	gql "github.com/shshimamo/knowledge/main/graph/model"
	"github.com/shshimamo/knowledge/main/middlewares"
	"github.com/shshimamo/knowledge/main/app/domain/user"
)

type UserUseCase interface {
	CreateUser(ctx context.Context, input *gql.NewUser) (*gql.User, error)
	GetUser(ctx context.Context, id int64) (*gql.User, error)
}

type userUseCase struct {
	userRepo user.UserRepository
}

func NewUserUseCase(userRepo user.UserRepository) UserUseCase {
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

	userModel := user.MapUserGqlNewToModel(gqlNew)
	userModel.AuthUserID = token.AuthUserID
	err := userModel.Validate()
	if err != nil {
		return nil, err
	}

	newUser, err := u.userRepo.CreateUser(ctx, userModel)
	if err != nil {
		return nil, err
	}

	gqlUser := user.MapUserModelToGql(newUser)

	return gqlUser, nil
}

func (u *userUseCase) GetUser(ctx context.Context, id int64) (*gql.User, error) {
	userModel, err := u.userRepo.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return user.MapUserModelToGql(userModel), nil
}