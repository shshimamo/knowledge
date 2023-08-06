package service

import (
	"context"
	"database/sql"
	gql "github.com/shshimamo/knowledge-main/graph/model"
	"github.com/shshimamo/knowledge-main/model"
	"github.com/shshimamo/knowledge-main/repository"
	"github.com/shshimamo/knowledge-main/service/util"
	"strconv"
)

type KnowledgeService interface {
	CreateKnowledge(ctx context.Context, gqlnew *gql.NewKnowledge) (*gql.Knowledge, error)
	GetKnowledge(ctx context.Context, id int) (*gql.Knowledge, error)
	GetMyKnowledge(ctx context.Context, id int) (*gql.Knowledge, error)
	UpdateKnowledge(ctx context.Context, id int, gqlupdate *gql.UpdateKnowledge) (*gql.Knowledge, error)
	DeleteKnowledge(ctx context.Context, id int) (*gql.DeleteKnowledgeResult, error)
}

type knowledgeService struct {
	db *sql.DB
}

func newKnowledgeService(db *sql.DB) *knowledgeService {
	return &knowledgeService{db}
}

func (s *knowledgeService) GetKnowledge(ctx context.Context, id int) (*gql.Knowledge, error) {
	repo := repository.NewKnowledgeRepository(s.db)
	k, err := repo.GetKnowledge(ctx, &repository.GetKnowledgeCommand{ID: id, IsPublic: true})
	if err != nil {
		return nil, err
	}

	gqlk := model.MapKnowledgeModelToGql(k)

	return gqlk, nil
}

func (s *knowledgeService) GetMyKnowledge(ctx context.Context, id int) (*gql.Knowledge, error) {
	my, err := util.CheckAuth(ctx)
	if err != nil {
		return nil, err
	}

	repo := repository.NewKnowledgeRepository(s.db)
	k, err := repo.GetKnowledge(ctx, &repository.GetKnowledgeCommand{ID: id, UserID: my.ID})
	if err != nil {
		return nil, err
	}

	gqlk := model.MapKnowledgeModelToGql(k)

	return gqlk, nil
}

func (s *knowledgeService) CreateKnowledge(ctx context.Context, gqlnew *gql.NewKnowledge) (*gql.Knowledge, error) {
	_, err := util.CheckAuth(ctx)
	if err != nil {
		return nil, err
	}

	k := model.MapKnowledgeGqlNewToModel(gqlnew)

	repo := repository.NewKnowledgeRepository(s.db)
	newk, err := repo.CreateKnowledge(ctx, k)
	if err != nil {
		return nil, err
	}

	gqlk := model.MapKnowledgeModelToGql(newk)

	return gqlk, nil
}

func (s *knowledgeService) UpdateKnowledge(ctx context.Context, id int, gqlupdate *gql.UpdateKnowledge) (*gql.Knowledge, error) {
	my, err := util.CheckAuth(ctx)
	if err != nil {
		return nil, err
	}

	repo := repository.NewKnowledgeRepository(s.db)
	k, err := repo.GetKnowledge(ctx, &repository.GetKnowledgeCommand{ID: id, UserID: my.ID})
	if err != nil {
		return nil, err
	}

	k.Title = gqlupdate.Title
	k.Text = gqlupdate.Text
	k.IsPublic = gqlupdate.IsPublic

	k2, err := repo.UpdateKnowledge(ctx, k)
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

	repo := repository.NewKnowledgeRepository(s.db)
	k, err := repo.GetKnowledge(ctx, &repository.GetKnowledgeCommand{ID: id, UserID: my.ID})
	if err != nil {
		return nil, err
	}

	err = repo.DeleteKnowledge(ctx, k)
	if err != nil {
		return nil, err
	}

	return &gql.DeleteKnowledgeResult{ID: strconv.Itoa(id)}, nil
}
