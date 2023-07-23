package repository

import (
	"context"

	"github.com/shshimamo/knowledge-main/model"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *model.User) (*model.User, error)
}

type userRepository struct {
	exec boil.ContextExecutor
}

func NewUserRepository(exec boil.ContextExecutor) UserRepository {
	return &userRepository{exec: exec}
}

func (r *userRepository) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	dbuser := model.MapUserModelToDB(user)

	err := dbuser.Insert(ctx, r.exec, boil.Infer())
	if err != nil {
		return nil, err
	}

	user.ID = dbuser.ID

	return user, nil
}
