package util

import (
	"context"
	"errors"

	"github.com/shshimamo/knowledge-main/middlewares/auth"
	"github.com/shshimamo/knowledge-main/model"
)

func CheckAuth(ctx context.Context) (*model.User, error) {
	user, ok := auth.GetCurrentUser(ctx)
	if !ok {
		return nil, errors.New("not authenticated")
	}
	return user, nil
}
