package util

import (
	"context"
	"github.com/shshimamo/knowledge/main/middlewares"
	"github.com/shshimamo/knowledge/main/model"
)

func CheckAuth(ctx context.Context) (*model.User, error) {
	user, ok := middlewares.GetCurrentUser(ctx)
	if !ok {
		return nil, ErrUnauthorized
	}
	return user, nil
}