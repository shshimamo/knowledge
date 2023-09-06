package repository

import (
	"context"
	"github.com/shshimamo/knowledge/auth/model"
)

type TokenRepository interface {
	Create(ctx context.Context, t *model.Token) (*model.Token, error)
}

type tokenRepository struct {
	dbtx DBTX
}

func NewTokenRepository(dbtx DBTX) TokenRepository {
	return &tokenRepository{dbtx: dbtx}
}

func (r *tokenRepository) Create(ctx context.Context, t *model.Token) (*model.Token, error) {
	err := r.dbtx.QueryRowContext(
		ctx,
		"INSERT INTO tokens (auth_user_id, token, active, expires_at) VALUES ($1, $2, $3, $4) RETURNING id",
		t.AuthUserID, t.Token, true, t.ExpiresAt,
	).Scan(&t.ID)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (r *tokenRepository) Deactivate(ctx context.Context, tokenStr string) error {
	_, err := r.dbtx.QueryContext(ctx, "UPDATE tokens SET active=$1 WHERE token=$2", false, tokenStr)
	if err != nil {
		return err
	}

	return nil
}
