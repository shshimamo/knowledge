package graph

import (
	"context"
	"errors"
	"github.com/shshimamo/knowledge/main/middlewares"

	"github.com/99designs/gqlgen/graphql"
	"github.com/shshimamo/knowledge/main/graph/generated"
)

var Directive generated.DirectiveRoot = generated.DirectiveRoot{
	IsAuthenticated: IsAuthenticated,
}

func IsAuthenticated(ctx context.Context, obj interface{}, next graphql.Resolver) (
	res interface{}, err error) {
	if _, ok := middlewares.GetCurrentUser(ctx); !ok {
		return nil, errors.New("not authenticated")
	}
	return next(ctx)
}
