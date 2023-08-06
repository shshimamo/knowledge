package model

import (
	"github.com/shshimamo/knowledge-main/db"
	gql "github.com/shshimamo/knowledge-main/graph/model"
	"strconv"
)

type Knowledge struct {
	ID       int
	UserID   int
	Title    string
	Text     string
	IsPublic bool
}

func MapKnowledgeDBToModel(db *db.Knowledge) *Knowledge {
	k := &Knowledge{
		ID:       int(db.ID),
		UserID:   int(db.UserID),
		Title:    db.Title,
		Text:     db.Text,
		IsPublic: db.IsPublic,
	}
	return k
}

func MapKnowledgeGqlNewToModel(gqlnew *gql.NewKnowledge) *Knowledge {
	k := &Knowledge{}
	k.Title = gqlnew.Title
	k.Text = gqlnew.Text
	k.IsPublic = gqlnew.IsPublic
	return k
}

func MapKnowledgeGqlUpdateToModel(gqlnew *gql.UpdateKnowledge) *Knowledge {
	k := &Knowledge{}
	k.Title = gqlnew.Title
	k.Text = gqlnew.Text
	k.IsPublic = gqlnew.IsPublic
	return k
}

func MapKnowledgeModelToDB(k *Knowledge) *db.Knowledge {
	db := &db.Knowledge{
		UserID:   int64(k.UserID),
		Title:    k.Title,
		Text:     k.Text,
		IsPublic: k.IsPublic,
	}
	return db
}

func MapKnowledgeModelToGql(k *Knowledge) *gql.Knowledge {
	gql := &gql.Knowledge{
		ID:       strconv.Itoa(k.ID),
		UserID:   strconv.Itoa(k.UserID),
		Title:    k.Title,
		Text:     k.Text,
		IsPublic: k.IsPublic,
	}
	return gql
}
