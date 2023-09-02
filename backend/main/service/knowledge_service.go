package service

import (
	"context"
	gql "github.com/shshimamo/knowledge-main/graph/model"
	"github.com/shshimamo/knowledge-main/model"
	"github.com/shshimamo/knowledge-main/repository"
	"github.com/shshimamo/knowledge-main/service/util"
	"strconv"
)

type KnowledgeService interface {
	CreateKnowledge(ctx context.Context, input *gql.CreateKnowledgeInput) (*gql.Knowledge, error)
	GetKnowledge(ctx context.Context, id int) (*gql.Knowledge, error)
	GetKnowledgeList(ctx context.Context, ids []int, uids []int) ([]*gql.Knowledge, error)
	GetMyKnowledge(ctx context.Context, id int) (*gql.Knowledge, error)
	UpdateKnowledge(ctx context.Context, id int, gqlupdate *gql.UpdateKnowledgeInput) (*gql.Knowledge, error)
	DeleteKnowledge(ctx context.Context, id int) (*gql.DeleteKnowledgeResult, error)
}

type knowledgeService struct {
	knowRepo repository.KnowledgeRepository
}

func newKnowledgeService(knowRepo repository.KnowledgeRepository) *knowledgeService {
	return &knowledgeService{knowRepo: knowRepo}
}

func (s *knowledgeService) GetKnowledge(ctx context.Context, id int) (*gql.Knowledge, error) {
	k, err := s.knowRepo.GetKnowledge(ctx, &repository.GetKnowledgeCommand{ID: id})
	if err != nil {
		return nil, err
	}

	if k.IsPublic {
		return model.MapKnowledgeModelToGql(k), nil
	}

	my, _ := util.CheckAuth(ctx)
	if my == nil || k.UserID != my.ID {
		return nil, util.ErrForbidden
	}

	return model.MapKnowledgeModelToGql(k), nil
}

func (s *knowledgeService) GetKnowledgeList(ctx context.Context, ids []int, uids []int) ([]*gql.Knowledge, error) {
	klist, err := s.knowRepo.GetKnowledgeList(ctx, &repository.GetKnowledgeListCommand{IDs: ids, UserIDs: uids})
	if err != nil {
		return nil, err
	}

	gqllist := model.MapKnowledgeListModelToGql(klist)

	return gqllist, nil
}

func (s *knowledgeService) GetMyKnowledge(ctx context.Context, id int) (*gql.Knowledge, error) {
	my, err := util.CheckAuth(ctx)
	if err != nil {
		return nil, err
	}

	k, err := s.knowRepo.GetKnowledge(ctx, &repository.GetKnowledgeCommand{ID: id, UserID: my.ID})
	if err != nil {
		return nil, err
	}

	gqlk := model.MapKnowledgeModelToGql(k)

	return gqlk, nil
}

func (s *knowledgeService) CreateKnowledge(ctx context.Context, input *gql.CreateKnowledgeInput) (*gql.Knowledge, error) {
	user, err := util.CheckAuth(ctx)
	if err != nil {
		return nil, err
	}

	k := model.MapKnowledgeGqlCreateInputToModel(input)
	k.UserID = user.ID

	newk, err := s.knowRepo.CreateKnowledge(ctx, k)
	if err != nil {
		return nil, err
	}

	gqlk := model.MapKnowledgeModelToGql(newk)

	return gqlk, nil
}

func (s *knowledgeService) UpdateKnowledge(ctx context.Context, id int, gqlupdate *gql.UpdateKnowledgeInput) (*gql.Knowledge, error) {
	my, err := util.CheckAuth(ctx)
	if err != nil {
		return nil, err
	}

	k, err := s.knowRepo.GetKnowledge(ctx, &repository.GetKnowledgeCommand{ID: id, UserID: my.ID})
	if err != nil {
		return nil, err
	}

	k.Title = gqlupdate.Title
	k.Text = gqlupdate.Text
	k.IsPublic = gqlupdate.IsPublic

	k2, err := s.knowRepo.UpdateKnowledge(ctx, k)
	if err != nil {
		return nil, err
	}

	gqlk := model.MapKnowledgeModelToGql(k2)
	return gqlk, nil
}

func (s *knowledgeService) DeleteKnowledge(ctx context.Context, id int) (*gql.DeleteKnowledgeResult, error) {
	my, err := util.CheckAuth(ctx)
	if err != nil {
		return nil, err
	}

	k, err := s.knowRepo.GetKnowledge(ctx, &repository.GetKnowledgeCommand{ID: id, UserID: my.ID})
	if err != nil {
		return nil, err
	}

	err = s.knowRepo.DeleteKnowledge(ctx, k)
	if err != nil {
		return nil, err
	}

	return &gql.DeleteKnowledgeResult{ID: strconv.Itoa(id)}, nil
}
