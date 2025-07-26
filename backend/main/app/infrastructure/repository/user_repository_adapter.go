package repository

import (
	"context"
	"database/sql"

	"github.com/aarondl/null/v8"
	"github.com/aarondl/sqlboiler/v4/boil"
	"github.com/shshimamo/knowledge/main/app/domain/user"
	"github.com/shshimamo/knowledge/main/db"
	"github.com/shshimamo/knowledge/main/model"
	"github.com/shshimamo/knowledge/main/repository"
)

// userRepositoryAdapter adapts the new domain repository to the old interface
type userRepositoryAdapter struct {
	domainRepo user.UserRepository
	exec       boil.ContextExecutor
}

func NewUserRepositoryAdapter(domainRepo user.UserRepository, exec boil.ContextExecutor) repository.UserRepository {
	return &userRepositoryAdapter{domainRepo: domainRepo, exec: exec}
}

func (a *userRepositoryAdapter) CreateUser(ctx context.Context, userModel *model.User) (*model.User, error) {
	// Convert model.User to domain user.User
	domainUser := &user.User{
		ID:         userModel.ID,
		AuthUserID: userModel.AuthUserID,
		Name:       userModel.Name,
	}

	newDomainUser, err := a.domainRepo.CreateUser(ctx, domainUser)
	if err != nil {
		return nil, err
	}

	// Convert back to model.User
	return &model.User{
		ID:         newDomainUser.ID,
		AuthUserID: newDomainUser.AuthUserID,
		Name:       newDomainUser.Name,
	}, nil
}

func (a *userRepositoryAdapter) GetUserByToken(ctx context.Context, token *model.Token) (*model.User, error) {
	dbUser, err := db.Users(db.UserWhere.AuthUserID.EQ(null.Int64From(int64(token.AuthUserID)))).One(ctx, a.exec)

	if err != nil {
		if err != sql.ErrNoRows {
			return nil, err
		} else {
			return nil, nil
		}
	}

	return model.MapUserDBToModel(dbUser), nil
}

func (a *userRepositoryAdapter) GetUser(ctx context.Context, cmd *repository.GetUserCommand) (*model.User, error) {
	var domainUser *user.User
	var err error

	if cmd.ID != 0 {
		domainUser, err = a.domainRepo.GetUserByID(ctx, cmd.ID)
	} else if cmd.Name != "" {
		domainUser, err = a.domainRepo.GetUserByUserName(ctx, cmd.Name)
	}

	if err != nil {
		return nil, err
	}

	return &model.User{
		ID:         domainUser.ID,
		AuthUserID: domainUser.AuthUserID,
		Name:       domainUser.Name,
	}, nil
}