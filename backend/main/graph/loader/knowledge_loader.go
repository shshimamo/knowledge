package loader

import (
	"context"
	"errors"
	"github.com/graph-gophers/dataloader/v7"
	gql "github.com/shshimamo/knowledge/main/graph/model"
	"github.com/shshimamo/knowledge/main/service"
	"github.com/shshimamo/knowledge/main/utils"
)

type knowledgeBatch struct {
	Srv service.AllService
}

func (u *knowledgeBatch) BatchGetKnowledgeList(ctx context.Context, userIDs []string) []*dataloader.Result[[]*gql.Knowledge] {
	// Initialize slice of results
	results := make([]*dataloader.Result[[]*gql.Knowledge], len(userIDs))
	for i := range results {
		// Initialize each result by setting an error
		results[i] = &dataloader.Result[[]*gql.Knowledge]{
			Error: errors.New("not found"),
		}
	}

	uids, err := utils.StringSliceToInt64Slice(userIDs)
	if err != nil {
		return results
	}

	klist, err := u.Srv.GetKnowledgeList(ctx, nil, uids)
	if err != nil {
		for i := range results {
			results[i] = &dataloader.Result[[]*gql.Knowledge]{Error: err}
		}
		return results
	}

	// each user slice of knowledge
	kMap := make(map[string][]*gql.Knowledge)
	for _, k := range klist {
		kMap[k.UserID] = append(kMap[k.UserID], k)
	}

	for i, userID := range userIDs {
		results[i] = &dataloader.Result[[]*gql.Knowledge]{Data: kMap[userID]}
	}

	return results
}
