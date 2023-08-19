package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	"github.com/rs/cors"
	"github.com/shshimamo/knowledge-auth/handler"
	"github.com/shshimamo/knowledge-auth/model"
)

type databaseConfig struct {
	host     string
	port     string
	user     string
	password string
	dbname   string
}

func main() {
	var err error

	appEnv := model.AppEnv(os.Getenv("APP_ENV"))

	db, err := setupDatabase(appEnv)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	port := getPort()
	h := setupHandler(db, appEnv)
	log.Fatal(http.ListenAndServe(":"+port, h))
}

func withContext(fn func(context.Context, http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		// set context using context.WithTimeout, context.WithDeadline, context.WithCancel
		fn(ctx, w, r)
	}
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
			dbname:   "knowledge-auth",
		}
	}
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbCfg.host, dbCfg.port, dbCfg.user, dbCfg.password, dbCfg.dbname)
	db, err := sql.Open("postgres", connStr)

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, err
}

func setupHandler(db *sql.DB, appEnv model.AppEnv) http.Handler {
	auth := handler.NewAuthHandler(db, appEnv)

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})
	mux.HandleFunc("/signup", withContext(auth.Signup))
	mux.HandleFunc("/signin", withContext(auth.Signin))
	mux.HandleFunc("/signout", withContext(auth.Signout))

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

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		return "8081"
	} else {
		return port
	}
}
