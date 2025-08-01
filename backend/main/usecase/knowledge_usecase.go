package usecase

import (
	"context"
	gql "github.com/shshimamo/knowledge/main/graph/model"
	"github.com/shshimamo/knowledge/main/model"
	"github.com/shshimamo/knowledge/main/repository"
	"github.com/shshimamo/knowledge/main/shared/util"
	"strconv"
)

type KnowledgeUseCase interface {
	CreateKnowledge(ctx context.Context, input *gql.CreateKnowledgeInput) (*gql.Knowledge, error)
	GetKnowledge(ctx context.Context, id int64) (*gql.Knowledge, error)
	GetKnowledgeList(ctx context.Context, ids []int64, uids []int64) ([]*gql.Knowledge, error)
	GetMyKnowledge(ctx context.Context, id int64) (*gql.Knowledge, error)
	UpdateKnowledge(ctx context.Context, id int64, gqlupdate *gql.UpdateKnowledgeInput) (*gql.Knowledge, error)
	DeleteKnowledge(ctx context.Context, id int64) (*gql.DeleteKnowledgeResult, error)
}

type knowledgeUseCase struct {
	knowRepo repository.KnowledgeRepository
}

func NewKnowledgeUseCase(knowRepo repository.KnowledgeRepository) KnowledgeUseCase {
	return &knowledgeUseCase{knowRepo: knowRepo}
}

func (u *knowledgeUseCase) GetKnowledge(ctx context.Context, id int64) (*gql.Knowledge, error) {
	k, err := u.knowRepo.GetKnowledge(ctx, &repository.GetKnowledgeCommand{ID: id})
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

func (u *knowledgeUseCase) GetKnowledgeList(ctx context.Context, ids []int64, uids []int64) ([]*gql.Knowledge, error) {
	klist, err := u.knowRepo.GetKnowledgeList(ctx, &repository.GetKnowledgeListCommand{IDs: ids, UserIDs: uids})
	if err != nil {
		return nil, err
	}

	gqllist := model.MapKnowledgeListModelToGql(klist)

	return gqllist, nil
}

func (u *knowledgeUseCase) GetMyKnowledge(ctx context.Context, id int64) (*gql.Knowledge, error) {
	my, err := util.CheckAuth(ctx)
	if err != nil {
		return nil, err
	}

	k, err := u.knowRepo.GetKnowledge(ctx, &repository.GetKnowledgeCommand{ID: id, UserID: my.ID})
	if err != nil {
		return nil, err
	}

	gqlk := model.MapKnowledgeModelToGql(k)

	return gqlk, nil
}

func (u *knowledgeUseCase) CreateKnowledge(ctx context.Context, input *gql.CreateKnowledgeInput) (*gql.Knowledge, error) {
	user, err := util.CheckAuth(ctx)
	if err != nil {
		return nil, err
	}

	k := model.MapKnowledgeGqlCreateInputToModel(input)
	k.UserID = user.ID

	newk, err := u.knowRepo.CreateKnowledge(ctx, k)
	if err != nil {
		return nil, err
	}

	gqlk := model.MapKnowledgeModelToGql(newk)

	return gqlk, nil
}

func (u *knowledgeUseCase) UpdateKnowledge(ctx context.Context, id int64, gqlupdate *gql.UpdateKnowledgeInput) (*gql.Knowledge, error) {
	my, err := util.CheckAuth(ctx)
	if err != nil {
		return nil, err
	}

	k, err := u.knowRepo.GetKnowledge(ctx, &repository.GetKnowledgeCommand{ID: id, UserID: my.ID})
	if err != nil {
		return nil, err
	}

	k.Title = gqlupdate.Title
	k.Text = gqlupdate.Text
	k.IsPublic = gqlupdate.IsPublic

	k2, err := u.knowRepo.UpdateKnowledge(ctx, k)
	if err != nil {
		return nil, err
	}

	gqlk := model.MapKnowledgeModelToGql(k2)
	return gqlk, nil
}

func (u *knowledgeUseCase) DeleteKnowledge(ctx context.Context, id int64) (*gql.DeleteKnowledgeResult, error) {
	my, err := util.CheckAuth(ctx)
	if err != nil {
		return nil, err
	}

	k, err := u.knowRepo.GetKnowledge(ctx, &repository.GetKnowledgeCommand{ID: id, UserID: my.ID})
	if err != nil {
		return nil, err
	}

	err = u.knowRepo.DeleteKnowledge(ctx, k)
	if err != nil {
		return nil, err
	}

	return &gql.DeleteKnowledgeResult{ID: strconv.FormatInt(id, 10)}, nil
}