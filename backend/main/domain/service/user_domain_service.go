package service

import (
	"context"
	"github.com/shshimamo/knowledge/main/model"
)

type UserDomainService interface {
	ValidateUser(ctx context.Context, user *model.User) error
	CanCreateUser(ctx context.Context, authUserID string) (bool, error)
}

type userDomainService struct{}

func NewUserDomainService() UserDomainService {
	return &userDomainService{}
}

func (d *userDomainService) ValidateUser(ctx context.Context, user *model.User) error {
	return user.Validate()
}

func (d *userDomainService) CanCreateUser(ctx context.Context, authUserID string) (bool, error) {
	return true, nil
}