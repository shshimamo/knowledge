package auth

import (
	"context"
	"database/sql"
	"github.com/shshimamo/knowledge-main/db"
	"github.com/shshimamo/knowledge-main/model"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"net/http"
)

type CurrentUserKey struct{}
type CurrentTokenKey struct{}

func NewAuthMiddleware(exec boil.ContextExecutor) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenStr := r.Header.Get("Authorization")

			if tokenStr == "" {
				next.ServeHTTP(w, r)
				return
			}

			token, err := model.NewToken(tokenStr)
			if err != nil {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}
			ctx := context.WithValue(r.Context(), CurrentTokenKey{}, token)

			dbUser, err := db.Users(db.UserWhere.AuthUserID.EQ(null.Int64From(int64(token.AuthUserID)))).One(r.Context(), exec)
			if err != nil {
				if err != sql.ErrNoRows {
					http.Error(w, "Forbidden", http.StatusForbidden)
					return
				}
			} else {
				user := model.MapUserDBToModel(dbUser)
				ctx = context.WithValue(r.Context(), CurrentUserKey{}, user)
			}

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func GetCurrentUser(ctx context.Context) (*model.User, bool) {
	switch v := ctx.Value(CurrentUserKey{}).(type) {
	case *model.User:
		return v, true
	default:
		return nil, false
	}
}

func GetCurrentToken(ctx context.Context) (*model.Token, bool) {
	switch v := ctx.Value(CurrentTokenKey{}).(type) {
	case *model.Token:
		return v, true
	default:
		return nil, false
	}
}