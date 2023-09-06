package main

import (
	"database/sql"
	"github.com/shshimamo/knowledge/auth/utils"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	"github.com/rs/cors"
	"github.com/shshimamo/knowledge/auth/handler"
	"github.com/shshimamo/knowledge/auth/model"
)

func main() {
	var err error

	appEnv := model.NewAppEnv()

	db, err := utils.SetupDatabase(appEnv)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	port := getPort()
	h := setupHandler(db, appEnv)
	log.Fatal(http.ListenAndServe(":"+port, h))
}

func setupHandler(db *sql.DB, appEnv model.AppEnv) http.Handler {
	auth := handler.NewAuthHandler(db, appEnv)

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("ok"))
	})
	mux.HandleFunc("/signup", auth.Signup)
	mux.HandleFunc("/signin", auth.Signin)
	mux.HandleFunc("/signout", auth.Signout)

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
