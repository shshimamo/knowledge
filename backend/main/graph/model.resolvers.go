package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.35

import (
	"context"

	"github.com/shshimamo/knowledge-main/graph/generated"
	"github.com/shshimamo/knowledge-main/graph/model"
)

// KnowledgeList is the resolver for the knowledgeList field.
func (r *currentUserResolver) KnowledgeList(ctx context.Context, obj *model.CurrentUser, first int) ([]*model.Knowledge, error) {
	thunk := r.Loaders.KnowledgeLoader.Load(ctx, obj.ID)
	klist, err := thunk()
	if err != nil {
		return nil, err
	}
	return klist, nil
}

// CurrentUser returns generated.CurrentUserResolver implementation.
func (r *Resolver) CurrentUser() generated.CurrentUserResolver { return &currentUserResolver{r} }

type currentUserResolver struct{ *Resolver }
