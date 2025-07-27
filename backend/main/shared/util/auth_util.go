package util

import (
	"context"
	"github.com/shshimamo/knowledge/main/app/presentation/middleware"
	"github.com/shshimamo/knowledge/main/model"
)

func CheckAuth(ctx context.Context) (*model.User, error) {
	user, ok := middleware.GetCurrentUser(ctx)
	if !ok {
		return nil, ErrUnauthorized
	}
	return user, nil
}