//go:generate mockgen -source=$GOFILE -package=mock$GOPACKAGE -destination=../mock/$GOPACKAGE/mock_$GOFILE
package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/shshimamo/knowledge/main/db"
	"github.com/shshimamo/knowledge/main/model"
	"github.com/shshimamo/knowledge/main/repository/errs"
	"github.com/volatiletech/null/v8"
	"github.com/aarondl/sqlboiler/v4/boil"
	"github.com/aarondl/sqlboiler/v4/queries/qm"
)

var _ UserRepository = (*userRepository)(nil)

type UserRepository interface {
	CreateUser(ctx context.Context, user *model.User) (*model.User, error)
	GetUserByToken(ctx context.Context, token *model.Token) (*model.User, error)
	GetUser(ctx context.Context, cmd *GetUserCommand) (*model.User, error)
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

type GetUserCommand struct {
	ID   int64
	Name string
}

func (r *userRepository) GetUser(ctx context.Context, cmd *GetUserCommand) (*model.User, error) {
	if cmd.ID == 0 && cmd.Name == "" {
		return nil, errors.New("id or name is required")
	}

	queryMods := make([]qm.QueryMod, 0)

	if cmd.ID != 0 {
		queryMods = append(queryMods, db.UserWhere.ID.EQ(cmd.ID))
	}
	if cmd.Name != "" {
		queryMods = append(queryMods, db.UserWhere.Name.EQ(null.StringFrom(cmd.Name)))
	}

	dbUser, err := db.Users(queryMods...).One(ctx, r.exec)
	if err != nil {
		return nil, errs.ConvertSqlError(err)
	}

	u := model.MapUserDBToModel(dbUser)

	return u, nil
}
