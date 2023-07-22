package auth

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/shshimamo/knowledge-main/db"
	"github.com/shshimamo/knowledge-main/model"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"net/http"
	"strconv"
)

type CurrentUserKey struct{}

type claims struct {
	AuthUserID string `json:"authUserId"`
	jwt.StandardClaims
}

func (c *claims) authUserIDInt() (int64, error) {
	num, err := strconv.ParseInt(c.AuthUserID, 10, 64)
	if err != nil {
		return 0, err
	}
	return num, nil
}

func getClaims(tokenStr string) (*claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &claims{}, func(token *jwt.Token) (interface{}, error) {
		// Always make sure the token method corresponds to the one you expect.
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// Return the key for validation
		return []byte("secret"), nil
	})

	claims, ok := token.Claims.(*claims)
	if ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

func NewAuthMiddleware(exec boil.ContextExecutor) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenStr := r.Header.Get("Authorization")

			if tokenStr == "" {
				next.ServeHTTP(w, r)
				return
			}

			claims, err := getClaims(tokenStr)

			if err != nil {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}

			auserID, err := claims.authUserIDInt()
			if err != nil {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}

			dbUser, err := db.Users(db.UserWhere.AuthUserID.EQ(null.Int64From(auserID))).One(r.Context(), exec)
			if err != nil {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}

			user := model.ConvertUserFromDB(dbUser)

			ctx := context.WithValue(r.Context(), CurrentUserKey{}, user)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
