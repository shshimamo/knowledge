package repository

import (
	"context"
	"errors"
	"github.com/shshimamo/knowledge/main/db"
	"github.com/shshimamo/knowledge/main/model"
	"github.com/shshimamo/knowledge/main/repository/errs"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

var _ KnowledgeRepository = (*knowledgeRepository)(nil)

type KnowledgeRepository interface {
	GetKnowledge(ctx context.Context, cmd *GetKnowledgeCommand) (*model.Knowledge, error)
	GetKnowledgeList(ctx context.Context, cmd *GetKnowledgeListCommand) ([]*model.Knowledge, error)
	CreateKnowledge(ctx context.Context, k *model.Knowledge) (*model.Knowledge, error)
	UpdateKnowledge(ctx context.Context, k *model.Knowledge) (*model.Knowledge, error)
	DeleteKnowledge(ctx context.Context, k *model.Knowledge) error
}

type knowledgeRepository struct {
	exec boil.ContextExecutor
}

func NewKnowledgeRepository(exec boil.ContextExecutor) KnowledgeRepository {
	return &knowledgeRepository{exec: exec}
}

type GetKnowledgeCommand struct {
	ID       int64
	UserID   int64
	IsPublic bool
}

func (r *knowledgeRepository) GetKnowledge(ctx context.Context, cmd *GetKnowledgeCommand) (*model.Knowledge, error) {
	if cmd.ID == 0 && cmd.UserID == 0 {
		return nil, errors.New("id or userid is required")
	}

	queryMods := make([]qm.QueryMod, 0)

	if cmd.ID != 0 {
		queryMods = append(queryMods, db.KnowledgeWhere.ID.EQ(int64(cmd.ID)))
	}
	if cmd.UserID != 0 {
		queryMods = append(queryMods, db.KnowledgeWhere.UserID.EQ(int64(cmd.UserID)))
	}
	if cmd.IsPublic {
		queryMods = append(queryMods, db.KnowledgeWhere.IsPublic.EQ(true))
	}

	dbk, err := db.Knowledges(queryMods...).One(ctx, r.exec)
	if err != nil {
		return nil, errs.ConvertSqlError(err)
	}

	k := model.MapKnowledgeDBToModel(dbk)

	return k, nil
}

type GetKnowledgeListCommand struct {
	IDs     []int64
	UserIDs []int64
}

func (r *knowledgeRepository) GetKnowledgeList(ctx context.Context, cmd *GetKnowledgeListCommand) ([]*model.Knowledge, error) {
	if len(cmd.IDs) == 0 && len(cmd.UserIDs) == 0 {
		return nil, errors.New("invalid command")
	}

	queryMods := make([]qm.QueryMod, 0)

	if len(cmd.IDs) > 0 {
		queryMods = append(queryMods, db.KnowledgeWhere.ID.IN(cmd.IDs))
	}

	if len(cmd.UserIDs) > 0 {
		queryMods = append(queryMods, db.KnowledgeWhere.UserID.IN(cmd.UserIDs))
	}

	dblist, err := db.Knowledges(queryMods...).All(ctx, r.exec)
	if err != nil {
		return nil, errs.ConvertSqlError(err)
	}

	return model.MapKnowledgeListDBToModel(dblist), nil
}

func (r *knowledgeRepository) CreateKnowledge(ctx context.Context, k *model.Knowledge) (*model.Knowledge, error) {
	dbk := model.MapKnowledgeModelToDB(k)

	err := dbk.Insert(ctx, r.exec, boil.Infer())
	if err != nil {
		return nil, errs.ConvertSqlError(err)
	}

	k.ID = dbk.ID

	return k, nil
}

func (r *knowledgeRepository) UpdateKnowledge(ctx context.Context, k *model.Knowledge) (*model.Knowledge, error) {
	dbk := model.MapKnowledgeModelToDB(k)
	_, err := dbk.Update(ctx, r.exec, boil.Infer())
	if err != nil {
		return nil, errs.ConvertSqlError(err)
	}

	k2 := model.MapKnowledgeDBToModel(dbk)
	return k2, nil
}

func (r *knowledgeRepository) DeleteKnowledge(ctx context.Context, k *model.Knowledge) error {
	dbk := model.MapKnowledgeModelToDB(k)
	_, err := dbk.Delete(ctx, r.exec)
	if err != nil {
		return errs.ConvertSqlError(err)
	}

	return nil
}
