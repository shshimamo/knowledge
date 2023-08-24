package service

import (
	"database/sql"
	"github.com/shshimamo/knowledge-main/repository"
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
	userRepo := repository.NewUserRepository(db)
	knowRepo := repository.NewKnowledgeRepository(db)
	return &allService{
		userService:      newUserService(userRepo),
		knowledgeService: newKnowledgeService(knowRepo),
	}
}
