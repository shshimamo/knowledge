package graph

import (
	"github.com/shshimamo/knowledge/main/graph/loader"
	"github.com/shshimamo/knowledge/main/usecase"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	AllUseCase usecase.AllUseCase
	Loaders    *loader.Loaders
}
