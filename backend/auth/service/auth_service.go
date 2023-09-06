package service

import (
	"context"
	"database/sql"
	"github.com/shshimamo/knowledge/auth/model"
	"github.com/shshimamo/knowledge/auth/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Signup(ctx context.Context, email string, password string) (*model.Token, error)
	Signin(ctx context.Context, email string, password string) (*model.Token, error)
	Signout(ctx context.Context, tokenStr string) error
}

type authService struct {
	db *sql.DB
}

func NewAuthService(db *sql.DB) AuthService {
	return &authService{db: db}
}

func (u *authService) Signup(ctx context.Context, email string, password string) (*model.Token, error) {
	auser := model.NewAuthUser(email)
	err := auser.CreatePasswordDigest(password)
	if err != nil {
		return nil, err
	}

	tx, err := u.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	auser, err = repository.NewAuthUserRepository(tx).Create(ctx, auser)
	if err != nil {
		return nil, err
	}

	token := model.NewToken(auser)
	token, err = repository.NewTokenRepository(tx).Create(ctx, token)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (u *authService) Signin(ctx context.Context, email string, password string) (*model.Token, error) {
	aurepo := repository.NewAuthUserRepository(u.db)
	auser, err := aurepo.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	// Validate password
	err = bcrypt.CompareHashAndPassword([]byte(auser.PasswordDigest), []byte(password))
	if err != nil {
		return nil, err
	}

	token := model.NewToken(auser)
	trepo := repository.NewTokenRepository(u.db)
	token, err = trepo.Create(ctx, token)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (u *authService) Signout(ctx context.Context, tokenStr string) error {
	_, err := u.db.Query("UPDATE tokens SET active=$1 WHERE token=$2", false, tokenStr)
	if err != nil {
		return err
	}

	return nil
}
