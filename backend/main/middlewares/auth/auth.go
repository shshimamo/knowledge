package auth

import (
	"context"
	"github.com/shshimamo/knowledge-main/model"
	"github.com/shshimamo/knowledge-main/repository"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"net/http"
)

func NewAuthMiddleware(exec boil.ContextExecutor) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			tokenStr := r.Header.Get("Authorization")

			if tokenStr == "" {
				if cookie, err := r.Cookie("token"); err == nil {
					tokenStr = cookie.Value
				}
			}

			if tokenStr == "" {
				next.ServeHTTP(w, r)
				return
			}

			token, err := model.NewToken(tokenStr)
			if err != nil {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}
			ctx = context.WithValue(ctx, model.CurrentTokenKey{}, token)

			repo := repository.NewUserRepository(exec)
			user, err := repo.GetUserByToken(ctx, token)

			if err != nil {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}

			ctx = context.WithValue(ctx, model.CurrentUserKey{}, user)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
