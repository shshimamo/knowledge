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

func MapKnowledgeDBToModel(dbk *db.Knowledge) *Knowledge {
	k := &Knowledge{
		ID:       int(dbk.ID),
		UserID:   int(dbk.UserID),
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

func MapKnowledgeGqlNewToModel(gqlnew *gql.NewKnowledge) *Knowledge {
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

func MapKnowledgeListModelToGql(klist []*Knowledge) []*gql.Knowledge {
	gqllist := make([]*gql.Knowledge, len(klist))
	for i, k := range klist {
		gqllist[i] = MapKnowledgeModelToGql(k)
	}
	return gqllist
}
