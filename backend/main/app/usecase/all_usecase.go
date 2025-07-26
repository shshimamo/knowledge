package usecase

import (
	"github.com/shshimamo/knowledge/main/app/usecase/knowledge"
	"github.com/shshimamo/knowledge/main/app/usecase/user"
	knowledgeDomain "github.com/shshimamo/knowledge/main/app/domain/knowledge"
	userDomain "github.com/shshimamo/knowledge/main/app/domain/user"
)

type AllUseCase interface {
	user.UserUseCase
	knowledge.KnowledgeUseCase
}

type allUseCase struct {
	user.UserUseCase
	knowledge.KnowledgeUseCase
}

func NewAllUseCase(userRepo userDomain.UserRepository, knowRepo knowledgeDomain.KnowledgeRepository) AllUseCase {
	return &allUseCase{
		UserUseCase:      user.NewUserUseCase(userRepo),
		KnowledgeUseCase: knowledge.NewKnowledgeUseCase(knowRepo),
	}
}