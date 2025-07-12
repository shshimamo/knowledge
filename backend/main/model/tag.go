package model

import (
	gql "github.com/shshimamo/knowledge/main/graph/model"
	"strconv"
)

type Tag struct {
	ID   int64
	Name string
}

type KnowledgeTag struct {
	ID          int64
	KnowledgeID int64
	TagID       int64
}

func MapTagModelToGql(t *Tag) *gql.Tag {
	return &gql.Tag{
		ID:   strconv.FormatInt(t.ID, 10),
		Name: t.Name,
	}
}

func MapTagListModelToGql(tags []*Tag) []*gql.Tag {
	gqlTags := make([]*gql.Tag, len(tags))
	for i, tag := range tags {
		gqlTags[i] = MapTagModelToGql(tag)
	}
	return gqlTags
}