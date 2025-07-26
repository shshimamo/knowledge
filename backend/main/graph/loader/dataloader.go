package loader

import (
	"github.com/shshimamo/knowledge/main/usecase"

	"github.com/graph-gophers/dataloader/v7"
	gql "github.com/shshimamo/knowledge/main/graph/model"
)

type Loaders struct {
	KnowledgeLoader dataloader.Interface[string, []*gql.Knowledge]
}

func NewLoaders(useCase usecase.AllUseCase) *Loaders {
	kb := &knowledgeBatch{UseCase: useCase}

	return &Loaders{
		KnowledgeLoader: dataloader.NewBatchedLoader[string, []*gql.Knowledge](
			kb.BatchGetKnowledgeList,
			dataloader.WithCache[string, []*gql.Knowledge](&dataloader.NoCache[string, []*gql.Knowledge]{}),
		),
	}
}
