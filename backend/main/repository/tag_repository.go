package repository

import (
	"context"
	"database/sql"
	"github.com/shshimamo/knowledge/main/model"
	"github.com/shshimamo/knowledge/main/repository/errs"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

var _ TagRepository = (*tagRepository)(nil)

type TagRepository interface {
	GetAllTags(ctx context.Context) ([]*model.Tag, error)
	GetTagsByKnowledgeID(ctx context.Context, knowledgeID int64) ([]*model.Tag, error)
	CreateOrGetTag(ctx context.Context, name string) (*model.Tag, error)
	SetKnowledgeTags(ctx context.Context, knowledgeID int64, tagNames []string) error
}

type tagRepository struct {
	exec boil.ContextExecutor
}

func NewTagRepository(exec boil.ContextExecutor) TagRepository {
	return &tagRepository{exec: exec}
}

func (r *tagRepository) GetAllTags(ctx context.Context) ([]*model.Tag, error) {
	query := `SELECT id, name FROM tags ORDER BY name`
	rows, err := r.exec.QueryContext(ctx, query)
	if err != nil {
		return nil, errs.ConvertSqlError(err)
	}
	defer rows.Close()

	var tags []*model.Tag
	for rows.Next() {
		tag := &model.Tag{}
		if err := rows.Scan(&tag.ID, &tag.Name); err != nil {
			return nil, errs.ConvertSqlError(err)
		}
		tags = append(tags, tag)
	}

	return tags, nil
}

func (r *tagRepository) GetTagsByKnowledgeID(ctx context.Context, knowledgeID int64) ([]*model.Tag, error) {
	query := `
		SELECT t.id, t.name 
		FROM tags t 
		INNER JOIN knowledge_tags kt ON t.id = kt.tag_id 
		WHERE kt.knowledge_id = $1 
		ORDER BY t.name`
	rows, err := r.exec.QueryContext(ctx, query, knowledgeID)
	if err != nil {
		return nil, errs.ConvertSqlError(err)
	}
	defer rows.Close()

	var tags []*model.Tag
	for rows.Next() {
		tag := &model.Tag{}
		if err := rows.Scan(&tag.ID, &tag.Name); err != nil {
			return nil, errs.ConvertSqlError(err)
		}
		tags = append(tags, tag)
	}

	return tags, nil
}

func (r *tagRepository) CreateOrGetTag(ctx context.Context, name string) (*model.Tag, error) {
	// まず既存のタグを検索
	query := `SELECT id, name FROM tags WHERE name = $1`
	tag := &model.Tag{}
	err := r.exec.QueryRowContext(ctx, query, name).Scan(&tag.ID, &tag.Name)
	if err == nil {
		return tag, nil
	}
	if err != sql.ErrNoRows {
		return nil, errs.ConvertSqlError(err)
	}

	// 存在しない場合は新規作成
	insertQuery := `INSERT INTO tags (name) VALUES ($1) RETURNING id, name`
	err = r.exec.QueryRowContext(ctx, insertQuery, name).Scan(&tag.ID, &tag.Name)
	if err != nil {
		return nil, errs.ConvertSqlError(err)
	}

	return tag, nil
}

func (r *tagRepository) SetKnowledgeTags(ctx context.Context, knowledgeID int64, tagNames []string) error {
	// トランザクション内で実行
	tx, ok := r.exec.(*sql.Tx)
	if !ok {
		// execがDBインスタンスの場合、新しいトランザクションを開始
		if db, isDB := r.exec.(*sql.DB); isDB {
			newTx, err := db.BeginTx(ctx, nil)
			if err != nil {
				return errs.ConvertSqlError(err)
			}
			defer newTx.Rollback()
			
			if err := r.setKnowledgeTagsInTx(ctx, newTx, knowledgeID, tagNames); err != nil {
				return err
			}
			return newTx.Commit()
		}
		return errs.ConvertSqlError(sql.ErrTxDone)
	}

	return r.setKnowledgeTagsInTx(ctx, tx, knowledgeID, tagNames)
}

func (r *tagRepository) setKnowledgeTagsInTx(ctx context.Context, tx *sql.Tx, knowledgeID int64, tagNames []string) error {
	// 既存の関連を削除
	_, err := tx.ExecContext(ctx, `DELETE FROM knowledge_tags WHERE knowledge_id = $1`, knowledgeID)
	if err != nil {
		return errs.ConvertSqlError(err)
	}

	// 新しいタグの関連を追加
	for _, tagName := range tagNames {
		if tagName == "" {
			continue
		}

		// タグを作成または取得
		tag, err := r.createOrGetTagInTx(ctx, tx, tagName)
		if err != nil {
			return err
		}

		// knowledge_tagsに関連を追加
		_, err = tx.ExecContext(ctx, 
			`INSERT INTO knowledge_tags (knowledge_id, tag_id) VALUES ($1, $2)`, 
			knowledgeID, tag.ID)
		if err != nil {
			return errs.ConvertSqlError(err)
		}
	}

	return nil
}

func (r *tagRepository) createOrGetTagInTx(ctx context.Context, tx *sql.Tx, name string) (*model.Tag, error) {
	// まず既存のタグを検索
	query := `SELECT id, name FROM tags WHERE name = $1`
	tag := &model.Tag{}
	err := tx.QueryRowContext(ctx, query, name).Scan(&tag.ID, &tag.Name)
	if err == nil {
		return tag, nil
	}
	if err != sql.ErrNoRows {
		return nil, errs.ConvertSqlError(err)
	}

	// 存在しない場合は新規作成
	insertQuery := `INSERT INTO tags (name) VALUES ($1) RETURNING id, name`
	err = tx.QueryRowContext(ctx, insertQuery, name).Scan(&tag.ID, &tag.Name)
	if err != nil {
		return nil, errs.ConvertSqlError(err)
	}

	return tag, nil
}