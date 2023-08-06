package auth

import (
	"context"
	"github.com/shshimamo/knowledge-main/model"
	"github.com/shshimamo/knowledge-main/repository"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"net/http"
)

type currentUserKey struct{}

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

			ctx = context.WithValue(ctx, currentUserKey{}, user)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func GetCurrentUser(ctx context.Context) (*model.User, bool) {
	switch v := ctx.Value(currentUserKey{}).(type) {
	case *model.User:
		if v == nil {
			return nil, false
		}
		return v, true
	default:
		return nil, false
	}
}
