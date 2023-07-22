package service

import "database/sql"

type AllService interface {
	UserService
}

type allService struct {
	*userService
}

func NewAllService(db *sql.DB) AllService {
	return &allService{
		userService: NewUserService(db),
	}
}
