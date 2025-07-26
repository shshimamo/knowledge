package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/aarondl/null/v8"
	"github.com/aarondl/sqlboiler/v4/boil"
	"github.com/shshimamo/knowledge/main/db"
	"github.com/shshimamo/knowledge/main/app/domain/user"
	"github.com/shshimamo/knowledge/main/model"
	"github.com/shshimamo/knowledge/main/repository/errs"
)

var _ user.UserRepository = (*userRepository)(nil)

type userRepository struct {
	exec boil.ContextExecutor
}

func NewUserRepository(exec boil.ContextExecutor) user.UserRepository {
	return &userRepository{exec: exec}
}

func (r *userRepository) CreateUser(ctx context.Context, userModel *user.User) (*user.User, error) {
	dbuser := user.MapUserModelToDB(userModel)

	err := dbuser.Insert(ctx, r.exec, boil.Infer())
	if err != nil {
		return nil, err
	}

	userModel.ID = dbuser.ID

	return userModel, nil
}

func (r *userRepository) GetUserByToken(ctx context.Context, token *model.Token) (*user.User, error) {
	dbUser, err := db.Users(db.UserWhere.AuthUserID.EQ(null.Int64From(int64(token.AuthUserID)))).One(ctx, r.exec)

	if err != nil {
		if err != sql.ErrNoRows {
			return nil, err
		} else {
			return nil, nil
		}
	}

	return user.MapUserDBToModel(dbUser), nil
}

func (r *userRepository) GetUserByID(ctx context.Context, userID int64) (*user.User, error) {
	if userID == 0 {
		return nil, errors.New("userID is required")
	}

	dbUser, err := db.Users(db.UserWhere.ID.EQ(userID)).One(ctx, r.exec)
	if err != nil {
		return nil, errs.ConvertSqlError(err)
	}

	u := user.MapUserDBToModel(dbUser)

	return u, nil
}

func (r *userRepository) GetUserByUserName(ctx context.Context, userName string) (*user.User, error) {
	if userName == "" {
		return nil, errors.New("userName is required")
	}

	dbUser, err := db.Users(db.UserWhere.Name.EQ(null.StringFrom(userName))).One(ctx, r.exec)
	if err != nil {
		return nil, errs.ConvertSqlError(err)
	}

	u := user.MapUserDBToModel(dbUser)

	return u, nil
}

func (r *userRepository) UpdateUser(ctx context.Context, userModel *user.User) (*user.User, error) {
	dbuser := user.MapUserModelToDB(userModel)
	_, err := dbuser.Update(ctx, r.exec, boil.Infer())
	if err != nil {
		return nil, errs.ConvertSqlError(err)
	}

	u := user.MapUserDBToModel(dbuser)
	return u, nil
}

func (r *userRepository) DeleteUser(ctx context.Context, userModel *user.User) error {
	dbuser := user.MapUserModelToDB(userModel)
	_, err := dbuser.Delete(ctx, r.exec)
	if err != nil {
		return errs.ConvertSqlError(err)
	}

	return nil
}