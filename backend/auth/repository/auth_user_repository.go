package repository

import (
	"context"
	"github.com/shshimamo/knowledge/auth/model"
)

type AuthUserRepository interface {
	Create(ctx context.Context, auser *model.AuthUser) (*model.AuthUser, error)
	FindByEmail(ctx context.Context, email string) (*model.AuthUser, error)
}

type authUserRepository struct {
	dbtx DBTX
}

func NewAuthUserRepository(dbtx DBTX) AuthUserRepository {
	return &authUserRepository{dbtx: dbtx}
}

func (r *authUserRepository) Create(ctx context.Context, auser *model.AuthUser) (*model.AuthUser, error) {
	err := r.dbtx.QueryRowContext(
		ctx,
		"INSERT INTO auth_users (email, password_digest) VALUES ($1, $2) RETURNING id", auser.Email, auser.PasswordDigest,
	).Scan(&auser.ID)
	if err != nil {
		return nil, err
	}

	return auser, err
}

func (r *authUserRepository) FindByEmail(ctx context.Context, email string) (*model.AuthUser, error) {
	auser := &model.AuthUser{}
	err := r.dbtx.QueryRowContext(
		ctx,
		"SELECT id, email, password_digest FROM auth_users WHERE email=$1", email,
	).Scan(&auser.ID, &auser.Email, &auser.PasswordDigest)
	if err != nil {
		return nil, err
	}

	return auser, nil
}
