package server

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/aarondl/sqlboiler/v4/boil"
	"github.com/rs/cors"

	"github.com/shshimamo/knowledge/main/app/usecase"
	"github.com/shshimamo/knowledge/main/app/config"
	infraRepo "github.com/shshimamo/knowledge/main/app/infrastructure/repository"
	"github.com/shshimamo/knowledge/main/graph"
	"github.com/shshimamo/knowledge/main/graph/generated"
	"github.com/shshimamo/knowledge/main/graph/loader"
	hand "github.com/shshimamo/knowledge/main/handler"
	"github.com/shshimamo/knowledge/main/middlewares"
	"github.com/shshimamo/knowledge/main/model"
	"github.com/shshimamo/knowledge/main/utils"
)

func Start(cfg *config.Config) error {
	appEnv := model.NewAppEnv()

	db, err := utils.SetupDatabase(appEnv)
	if err != nil {
		return err
	}
	defer func() { _ = db.Close() }()

	sqlDebug := os.Getenv("SQL_DEBUG")
	if sqlDebug == "true" {
		boil.DebugMode = true
	}

	h := setupHandler(db, appEnv)
	port := cfg.Port
	if port == "" {
		port = "8080"
	}

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	return http.ListenAndServe(":"+port, h)
}

func setupHandler(exec boil.ContextExecutor, appEnv model.AppEnv) http.Handler {
	userDomainRepo := infraRepo.NewUserRepository(exec)
	userRepoAdapter := infraRepo.NewUserRepositoryAdapter(userDomainRepo, exec)
	
	allUseCase := usecase.NewAllUseCase(
		userDomainRepo,
		infraRepo.NewKnowledgeRepository(exec),
	)
	schema := generated.NewExecutableSchema(generated.Config{
		Resolvers: &graph.Resolver{
			AllUseCase: allUseCase,
			Loaders:    loader.NewLoaders(allUseCase),
		},
		Directives: graph.Directive,
	})
	srv := handler.New(schema)
	srv.AddTransport(transport.POST{})

	mux := http.NewServeMux()
	mux.Handle("/", playground.Handler("GraphQL playground", "/query"))
	mux.Handle("/query", middlewares.NewSlogMiddleware(middlewares.NewAuthMiddleware(userRepoAdapter)(srv)))

	th := hand.NewTokenHandler(appEnv)
	mux.Handle("/set_token", middlewares.NewSlogMiddleware(http.HandlerFunc(th.SetToken)))

	var allowOrigins []string
	if appEnv == model.Production {
		allowOrigins = []string{"http://frontend.main.shshimamo.com"}
	} else {
		allowOrigins = []string{"http://localhost:3000"}
	}
	c := cors.New(cors.Options{
		AllowedOrigins:   allowOrigins,
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "HEAD", "POST", "PUT", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
	})
	h := c.Handler(mux)

	return h
}