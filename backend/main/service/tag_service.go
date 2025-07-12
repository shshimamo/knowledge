package service

import (
	"context"
	gql "github.com/shshimamo/knowledge/main/graph/model"
	"github.com/shshimamo/knowledge/main/model"
	"github.com/shshimamo/knowledge/main/repository"
)

var _ TagService = (*tagService)(nil)

type TagService interface {
	GetAllTags(ctx context.Context) ([]*gql.Tag, error)
}

type tagService struct {
	tagRepo repository.TagRepository
}

func newTagService(tagRepo repository.TagRepository) *tagService {
	return &tagService{tagRepo: tagRepo}
}

func (s *tagService) GetAllTags(ctx context.Context) ([]*gql.Tag, error) {
	tags, err := s.tagRepo.GetAllTags(ctx)
	if err != nil {
		return nil, err
	}

	return model.MapTagListModelToGql(tags), nil
}