package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/shshimamo/knowledge/backend/oauth/internal/handler"
	"github.com/shshimamo/knowledge/backend/oauth/internal/util/bunapp"
	"github.com/shshimamo/knowledge/backend/oauth/internal/util/httputil"
)

func setRouter() {
	bunapp.OnStart("example.init", func(ctx context.Context, app *bunapp.App) error {
		router := app.Router()
		//api := app.APIRouter()

		welcomeHandler := handler.NewWelcomeHandler(app)

		router.GET("/", welcomeHandler.Welcome)

		return nil
	})
}

func main() {
	env := getEnv()
	if env != "dev" && env != "test" {
		log.Fatal("Invalid APP_ENV: must be either 'dev' or 'test'")
	}

	setRouter()
	ctx, app, err := bunapp.Start(context.Background(), "api", env)
	if err != nil {
		log.Fatal(err)
	}
	defer app.Stop()

	var handler http.Handler
	handler = app.Router()
	handler = httputil.ExitOnPanicHandler{Next: handler}

	srv := &http.Server{
		Addr:         "localhost:8010",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
		Handler:      handler,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && !isServerClosed(err) {
			log.Printf("ListenAndServe failed: %s", err)
		}
	}()

	fmt.Printf("listening on http://%s\n", srv.Addr)
	fmt.Println(bunapp.WaitExitSignal())

	err = srv.Shutdown(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

func isServerClosed(err error) bool {
	return err.Error() == "http: Server closed"
}

func getEnv() string {
	env := os.Getenv("APP_ENV")
	if env == "" {
		return "dev"
	}
	return env
}
