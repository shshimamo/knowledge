package main

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"github.com/shshimamo/knowledge-main/model"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
	"github.com/shshimamo/knowledge-main/graph"
	"github.com/shshimamo/knowledge-main/graph/generated"
	hand "github.com/shshimamo/knowledge-main/handler"
	"github.com/shshimamo/knowledge-main/middlewares/auth"
	"github.com/shshimamo/knowledge-main/service"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type databaseConfig struct {
	host     string
	port     string
	user     string
	password string
	dbname   string
}

func main() {
	appEnv := model.AppEnv(os.Getenv("APP_ENV"))

	db, err := setupDatabase(appEnv)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

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

func setupDatabase(env model.AppEnv) (*sql.DB, error) {
	var dbCfg databaseConfig
	if env == model.Production {
		dbCfg = databaseConfig{
			host:     os.Getenv("DB_HOST"),
			port:     os.Getenv("DB_PORT"),
			user:     os.Getenv("DB_USER"),
			password: os.Getenv("DB_PASSWORD"),
			dbname:   os.Getenv("DB_NAME"),
		}
	} else {
		dbCfg = databaseConfig{
			host:     "localhost",
			port:     "5432",
			user:     "postgres",
			password: "password",
			dbname:   "knowledge-main",
		}
	}
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbCfg.host, dbCfg.port, dbCfg.user, dbCfg.password, dbCfg.dbname)
	db, err := sql.Open("postgres", connStr)

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, err
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		return "8080"
	} else {
		return port
	}
}

func setupHandler(db *sql.DB, appEnv model.AppEnv) http.Handler {
	mux := http.NewServeMux()

	allService := service.NewAllService(db)
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: &graph.Resolver{
			AllService: allService,
			//Loaders: loader.NewLoaders(service),
		},
		//Directives: graph.Directive,
	}))
	//gqlMiddleware(srv)

	mux.Handle("/", playground.Handler("GraphQL playground", "/query"))
	mux.Handle("/query", auth.NewAuthMiddleware(db)(srv))

	th := hand.NewTokenHandler(appEnv)
	mux.HandleFunc("/set_token", withContext(th.SetToken))

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "HEAD", "POST", "PUT", "OPTIONS"},
		AllowedHeaders:   []string{"*"}, // Allow All HTTP Headers
	})
	h := c.Handler(mux)

	return h
}

func gqlMiddleware(srv *handler.Server) {
	srv.AroundOperations(func(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
		log.Println("before OperationHandler")
		res := next(ctx)
		defer log.Println("after OperationHandler")
		return res
	})
	srv.AroundResponses(func(ctx context.Context, next graphql.ResponseHandler) *graphql.Response {
		log.Println("before ResponseHandler")
		res := next(ctx)
		defer log.Println("after ResponseHandler")
		return res
	})
	srv.AroundRootFields(func(ctx context.Context, next graphql.RootResolver) graphql.Marshaler {
		log.Println("before RootResolver")
		res := next(ctx)
		defer func() {
			var b bytes.Buffer
			res.MarshalGQL(&b)
			log.Println("after RootResolver", b.String())
		}()
		return res
	})
	srv.AroundFields(func(ctx context.Context, next graphql.Resolver) (res interface{}, err error) {
		log.Println("before Resolver")
		res, err = next(ctx)
		defer log.Println("after Resolver", res)
		return
	})
}

func withContext(fn func(context.Context, http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		// set context using context.WithTimeout, context.WithDeadline, context.WithCancel
		fn(ctx, w, r)
	}
}
