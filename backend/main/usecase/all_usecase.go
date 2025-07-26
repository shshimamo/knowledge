package usecase

import (
	"github.com/shshimamo/knowledge/main/repository"
)

type AllUseCase interface {
	UserUseCase
	KnowledgeUseCase
}

type allUseCase struct {
	UserUseCase
	KnowledgeUseCase
}

func NewAllUseCase(userRepo repository.UserRepository, knowRepo repository.KnowledgeRepository) AllUseCase {
	return &allUseCase{
		UserUseCase:      NewUserUseCase(userRepo),
		KnowledgeUseCase: NewKnowledgeUseCase(knowRepo),
	}
}