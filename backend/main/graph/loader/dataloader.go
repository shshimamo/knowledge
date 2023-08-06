package loader

import (
	"github.com/shshimamo/knowledge-main/service"

	"github.com/graph-gophers/dataloader/v7"
	gql "github.com/shshimamo/knowledge-main/graph/model"
)

type Loaders struct {
	KnowledgeLoader dataloader.Interface[string, []*gql.Knowledge]
}

func NewLoaders(Srv service.AllService) *Loaders {
	kb := &knowledgeBatch{Srv: Srv}

	return &Loaders{
		KnowledgeLoader: dataloader.NewBatchedLoader[string, []*gql.Knowledge](
			kb.BatchGetKnowledgeList,
			dataloader.WithCache[string, []*gql.Knowledge](&dataloader.NoCache[string, []*gql.Knowledge]{}),
		),
	}
}
