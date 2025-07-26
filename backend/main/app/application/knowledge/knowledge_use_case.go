package knowledge

import (
	"context"
	"strconv"

	gql "github.com/shshimamo/knowledge/main/graph/model"
	"github.com/shshimamo/knowledge/main/app/domain/knowledge"
	"github.com/shshimamo/knowledge/main/shared/util"
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
	knowRepo knowledge.KnowledgeRepository
}

func NewKnowledgeUseCase(knowRepo knowledge.KnowledgeRepository) KnowledgeUseCase {
	return &knowledgeUseCase{knowRepo: knowRepo}
}

func (u *knowledgeUseCase) GetKnowledge(ctx context.Context, id int64) (*gql.Knowledge, error) {
	k, err := u.knowRepo.GetKnowledge(ctx, &knowledge.GetKnowledgeCommand{ID: id})
	if err != nil {
		return nil, err
	}

	if k.IsPublic {
		return knowledge.MapKnowledgeModelToGql(k), nil
	}

	my, _ := util.CheckAuth(ctx)
	if my == nil || k.UserID != my.ID {
		return nil, util.ErrForbidden
	}

	return knowledge.MapKnowledgeModelToGql(k), nil
}

func (u *knowledgeUseCase) GetKnowledgeList(ctx context.Context, ids []int64, uids []int64) ([]*gql.Knowledge, error) {
	klist, err := u.knowRepo.GetKnowledgeList(ctx, &knowledge.GetKnowledgeListCommand{IDs: ids, UserIDs: uids})
	if err != nil {
		return nil, err
	}

	gqllist := knowledge.MapKnowledgeListModelToGql(klist)

	return gqllist, nil
}

func (u *knowledgeUseCase) GetMyKnowledge(ctx context.Context, id int64) (*gql.Knowledge, error) {
	my, err := util.CheckAuth(ctx)
	if err != nil {
		return nil, err
	}

	k, err := u.knowRepo.GetKnowledge(ctx, &knowledge.GetKnowledgeCommand{ID: id, UserID: my.ID})
	if err != nil {
		return nil, err
	}

	gqlk := knowledge.MapKnowledgeModelToGql(k)

	return gqlk, nil
}

func (u *knowledgeUseCase) CreateKnowledge(ctx context.Context, input *gql.CreateKnowledgeInput) (*gql.Knowledge, error) {
	user, err := util.CheckAuth(ctx)
	if err != nil {
		return nil, err
	}

	k := knowledge.MapKnowledgeGqlCreateInputToModel(input)
	k.UserID = user.ID

	newk, err := u.knowRepo.CreateKnowledge(ctx, k)
	if err != nil {
		return nil, err
	}

	gqlk := knowledge.MapKnowledgeModelToGql(newk)

	return gqlk, nil
}

func (u *knowledgeUseCase) UpdateKnowledge(ctx context.Context, id int64, gqlupdate *gql.UpdateKnowledgeInput) (*gql.Knowledge, error) {
	my, err := util.CheckAuth(ctx)
	if err != nil {
		return nil, err
	}

	k, err := u.knowRepo.GetKnowledge(ctx, &knowledge.GetKnowledgeCommand{ID: id, UserID: my.ID})
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

	gqlk := knowledge.MapKnowledgeModelToGql(k2)
	return gqlk, nil
}

func (u *knowledgeUseCase) DeleteKnowledge(ctx context.Context, id int64) (*gql.DeleteKnowledgeResult, error) {
	my, err := util.CheckAuth(ctx)
	if err != nil {
		return nil, err
	}

	k, err := u.knowRepo.GetKnowledge(ctx, &knowledge.GetKnowledgeCommand{ID: id, UserID: my.ID})
	if err != nil {
		return nil, err
	}

	err = u.knowRepo.DeleteKnowledge(ctx, k)
	if err != nil {
		return nil, err
	}

	return &gql.DeleteKnowledgeResult{ID: strconv.FormatInt(id, 10)}, nil
}