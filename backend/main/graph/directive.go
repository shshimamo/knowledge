package graph

import (
	"context"
	"errors"

	"github.com/99designs/gqlgen/graphql"
	"github.com/shshimamo/knowledge-main/graph/generated"
	"github.com/shshimamo/knowledge-main/middlewares/auth"
)

var Directive generated.DirectiveRoot = generated.DirectiveRoot{
	IsAuthenticated: IsAuthenticated,
}

func IsAuthenticated(ctx context.Context, obj interface{}, next graphql.Resolver) (
	res interface{}, err error) {
	if _, ok := auth.GetCurrentUser(ctx); !ok {
		return nil, errors.New("not authenticated")
	}
	return next(ctx)
}
