package auth

import (
	"github.com/volatiletech/sqlboiler/v4/boil"
	"net/http"
)

func NewAuthCookieMiddleware(exec boil.ContextExecutor) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenStr := r.Header.Get("Authorization")

		})
	}
}
