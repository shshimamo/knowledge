package service

import (
	"context"
	"github.com/shshimamo/knowledge/main/model"
	"github.com/shshimamo/knowledge/main/shared/util"
)

type KnowledgeDomainService interface {
	CanAccessKnowledge(ctx context.Context, knowledge *model.Knowledge, userID int64) (bool, error)
	ValidateKnowledgeOwnership(ctx context.Context, knowledge *model.Knowledge, userID int64) error
}

type knowledgeDomainService struct{}

func NewKnowledgeDomainService() KnowledgeDomainService {
	return &knowledgeDomainService{}
}

func (d *knowledgeDomainService) CanAccessKnowledge(ctx context.Context, knowledge *model.Knowledge, userID int64) (bool, error) {
	if knowledge.IsPublic {
		return true, nil
	}
	
	return knowledge.UserID == userID, nil
}

func (d *knowledgeDomainService) ValidateKnowledgeOwnership(ctx context.Context, knowledge *model.Knowledge, userID int64) error {
	if knowledge.UserID != userID {
		return util.ErrForbidden
	}
	return nil
}