package model

import (
	"github.com/shshimamo/knowledge/main/db"
	gql "github.com/shshimamo/knowledge/main/graph/model"
	"strconv"
)

type Knowledge struct {
	ID       int64
	UserID   int64
	Title    string
	Text     string
	IsPublic bool
}

func MapKnowledgeDBToModel(dbk *db.Knowledge) *Knowledge {
	k := &Knowledge{
		ID:       dbk.ID,
		UserID:   dbk.UserID,
		Title:    dbk.Title,
		Text:     dbk.Text,
		IsPublic: dbk.IsPublic,
	}
	return k
}

func MapKnowledgeListDBToModel(dblist []*db.Knowledge) []*Knowledge {
	klist := make([]*Knowledge, len(dblist))
	for i, dbk := range dblist {
		klist[i] = MapKnowledgeDBToModel(dbk)
	}
	return klist
}

func MapKnowledgeGqlCreateInputToModel(input *gql.CreateKnowledgeInput) *Knowledge {
	k := &Knowledge{}
	k.Title = input.Title
	k.Text = input.Text
	k.IsPublic = input.IsPublic
	return k
}

func MapKnowledgeModelToDB(k *Knowledge) *db.Knowledge {
	db := &db.Knowledge{
		ID:       int64(k.ID),
		UserID:   int64(k.UserID),
		Title:    k.Title,
		Text:     k.Text,
		IsPublic: k.IsPublic,
	}
	return db
}

func MapKnowledgeModelToGql(k *Knowledge) *gql.Knowledge {
	gql := &gql.Knowledge{
		ID:       strconv.FormatInt(k.ID, 10),
		UserID:   strconv.FormatInt(k.UserID, 10),
		Title:    k.Title,
		Text:     k.Text,
		IsPublic: k.IsPublic,
	}
	return gql
}

func MapKnowledgeListModelToGql(klist []*Knowledge) []*gql.Knowledge {
	gqllist := make([]*gql.Knowledge, len(klist))
	for i, k := range klist {
		gqllist[i] = MapKnowledgeModelToGql(k)
	}
	return gqllist
}
