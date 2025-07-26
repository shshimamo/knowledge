package knowledge

import (
	"context"
)

type KnowledgeRepository interface {
	GetKnowledge(ctx context.Context, cmd *GetKnowledgeCommand) (*Knowledge, error)
	GetKnowledgeList(ctx context.Context, cmd *GetKnowledgeListCommand) ([]*Knowledge, error)
	CreateKnowledge(ctx context.Context, k *Knowledge) (*Knowledge, error)
	UpdateKnowledge(ctx context.Context, k *Knowledge) (*Knowledge, error)
	DeleteKnowledge(ctx context.Context, k *Knowledge) error
}

type GetKnowledgeCommand struct {
	ID       int64
	UserID   int64
	IsPublic bool
}

type GetKnowledgeListCommand struct {
	IDs     []int64
	UserIDs []int64
}