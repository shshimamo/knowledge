package util

import (
	"context"
	"github.com/shshimamo/knowledge-main/middlewares/auth"
	"github.com/shshimamo/knowledge-main/model"
)

func CheckAuth(ctx context.Context) (*model.User, error) {
	user, ok := auth.GetCurrentUser(ctx)
	if !ok {
		return nil, ErrUnauthorized
	}
	return user, nil
}
