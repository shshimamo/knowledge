package service

import (
	"github.com/shshimamo/knowledge/main/repository"
)

var _ AllService = (*allService)(nil)

type AllService interface {
	UserService
	KnowledgeService
	TagService
}

type allService struct {
	*userService
	*knowledgeService
	*tagService
}

func NewAllService(userRepo repository.UserRepository, knowRepo repository.KnowledgeRepository, tagRepo repository.TagRepository) AllService {
	return &allService{
		userService:      newUserService(userRepo),
		knowledgeService: newKnowledgeService(knowRepo, tagRepo),
		tagService:       newTagService(tagRepo),
	}
}
