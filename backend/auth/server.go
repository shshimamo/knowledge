package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	"github.com/shshimamo/knowledge-auth/handler"
)

type databaseConfig struct {
	host     string
	port     string
	user     string
	password string
	dbname   string
}

type AppEnv string

const (
	Production AppEnv = "production"
)

func main() {
	var err error

	appEnv := AppEnv(os.Getenv("APP_ENV"))

	db, err := setupDatabase(appEnv)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	h := handler.New(db)

	mux := http.NewServeMux()
	mux.HandleFunc("/signup", withContext(h.Signup))
	mux.HandleFunc("/signin", withContext(h.Signin))
	mux.HandleFunc("/signout", withContext(h.Signout))

	log.Fatal(http.ListenAndServe(":80", mux))
}

func withContext(fn func(context.Context, http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		// set context using context.WithTimeout, context.WithDeadline, context.WithCancel
		fn(ctx, w, r)
	}
}

func setupDatabase(env AppEnv) (*sql.DB, error) {
	var dbCfg databaseConfig
	if env == Production {
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
