package service

import (
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

func NewAllService(userRepo repository.UserRepository, knowRepo repository.KnowledgeRepository) AllService {
	return &allService{
		userService:      newUserService(userRepo),
		knowledgeService: newKnowledgeService(knowRepo),
	}
}
