package main

import (
	"context"
	"github.com/shshimamo/knowledge-main/graph/loader"
	"github.com/shshimamo/knowledge-main/middlewares"
	"github.com/shshimamo/knowledge-main/model"
	"github.com/shshimamo/knowledge-main/repository"
	"github.com/shshimamo/knowledge-main/utils"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/rs/cors"
	"github.com/shshimamo/knowledge-main/graph"
	"github.com/shshimamo/knowledge-main/graph/generated"
	hand "github.com/shshimamo/knowledge-main/handler"
	"github.com/shshimamo/knowledge-main/service"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func main() {
	appEnv := model.NewAppEnv()

	db, err := utils.SetupDatabase(appEnv)
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = db.Close() }()

	sqlDebug := os.Getenv("SQL_DEBUG")
	if sqlDebug == "true" {
		// Output SQLBoiler's SQL query log
		boil.DebugMode = true
	}

	h := setupHandler(db, appEnv)
	port := getPort()
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, h))
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		return "8080"
	} else {
		return port
	}
}

func setupHandler(exec boil.ContextExecutor, appEnv model.AppEnv) http.Handler {
	userRepo := repository.NewUserRepository(exec)
	allService := service.NewAllService(
		userRepo,
		repository.NewKnowledgeRepository(exec),
	)
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: &graph.Resolver{
			AllService: allService,
			Loaders:    loader.NewLoaders(allService),
		},
		Directives: graph.Directive,
	}))
	//gqlMiddleware(srv)

	mux := http.NewServeMux()
	mux.Handle("/", playground.Handler("GraphQL playground", "/query"))
	mux.Handle("/query", middlewares.NewAuthMiddleware(userRepo)(srv))

	th := hand.NewTokenHandler(appEnv)
	mux.HandleFunc("/set_token", withContext(th.SetToken))

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
		AllowedHeaders:   []string{"*"}, // Allow All HTTP Headers
	})
	h := c.Handler(mux)

	return h
}

//func gqlMiddleware(srv *handler.Server) {
//	srv.AroundOperations(func(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
//		log.Println("before OperationHandler")
//		res := next(ctx)
//		defer log.Println("after OperationHandler")
//		return res
//	})
//	srv.AroundResponses(func(ctx context.Context, next graphql.ResponseHandler) *graphql.Response {
//		log.Println("before ResponseHandler")
//		res := next(ctx)
//		defer log.Println("after ResponseHandler")
//		return res
//	})
//	srv.AroundRootFields(func(ctx context.Context, next graphql.RootResolver) graphql.Marshaler {
//		log.Println("before RootResolver")
//		res := next(ctx)
//		defer func() {
//			var b bytes.Buffer
//			res.MarshalGQL(&b)
//			log.Println("after RootResolver", b.String())
//		}()
//		return res
//	})
//	srv.AroundFields(func(ctx context.Context, next graphql.Resolver) (res interface{}, err error) {
//		log.Println("before Resolver")
//		res, err = next(ctx)
//		defer log.Println("after Resolver", res)
//		return
//	})
//}

func withContext(fn func(context.Context, http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		// set context using context.WithTimeout, context.WithDeadline, context.WithCancel
		fn(ctx, w, r)
	}
}
