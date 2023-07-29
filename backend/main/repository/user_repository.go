package repository

import (
	"context"
	"database/sql"
	"github.com/shshimamo/knowledge-main/db"
	"github.com/shshimamo/knowledge-main/model"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *model.User) (*model.User, error)
	GetUserByToken(ctx context.Context, token *model.Token) (*model.User, error)
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

func (r *userRepository) GetUserByToken(ctx context.Context, token *model.Token) (*model.User, error) {
	dbUser, err := db.Users(db.UserWhere.AuthUserID.EQ(null.Int64From(int64(token.AuthUserID)))).One(ctx, r.exec)

	if err != nil {
		if err != sql.ErrNoRows {
			return nil, err
		} else {
			return nil, nil
		}
	}

	return model.MapUserDBToModel(dbUser), nil
}
