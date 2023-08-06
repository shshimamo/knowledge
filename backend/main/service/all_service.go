package service

import (
	"database/sql"
)

type AllService interface {
	UserService
	KnowledgeService
}

type allService struct {
	*userService
	*knowledgeService
}

func NewAllService(db *sql.DB) AllService {
	return &allService{
		userService:      newUserService(db),
		knowledgeService: newKnowledgeService(db),
	}
}
