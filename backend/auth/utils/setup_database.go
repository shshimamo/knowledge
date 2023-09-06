package utils

import (
	"database/sql"
	"fmt"
	"github.com/shshimamo/knowledge/auth/model"
	"os"

	_ "github.com/lib/pq"
)

type databaseConfig struct {
	host     string
	port     string
	user     string
	password string
	dbname   string
}

func SetupDatabase(env model.AppEnv) (*sql.DB, error) {
	dbcfg := databaseConfig{
		host:     os.Getenv("DB_HOST"),
		port:     os.Getenv("DB_PORT"),
		user:     os.Getenv("DB_USER"),
		password: os.Getenv("DB_PASSWORD"),
		dbname:   "knowledge-auth-" + string(env),
	}

	if env == model.Develop || env == model.Test {
		if dbcfg.host == "" {
			dbcfg.host = "localhost"
		}
		if dbcfg.port == "" {
			dbcfg.port = "5432"
		}
		if dbcfg.user == "" {
			dbcfg.user = "postgres"
		}
		if dbcfg.password == "" {
			dbcfg.password = "password"
		}
	}

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbcfg.host, dbcfg.port, dbcfg.user, dbcfg.password, dbcfg.dbname)
	db, err := sql.Open("postgres", connStr)

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, err
}
