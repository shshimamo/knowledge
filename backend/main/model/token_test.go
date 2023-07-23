package model

import (
	"github.com/dgrijalva/jwt-go"
	"strings"
	"testing"
	"time"
)

func TestGetClaims(t *testing.T) {
	t.Parallel()

	validSecret := "secret" // TODO: ENV

	// valid token
	validClaims := &claims{
		AuthUserID: "12345",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 7).Unix(),
		},
	}
	validToken := jwt.NewWithClaims(jwt.SigningMethodHS256, validClaims)
	validTokenStr, _ := validToken.SignedString([]byte(validSecret))

	// invalid claims
	invalidClaimsToken := jwt.New(jwt.SigningMethodHS256)
	invalidClaimsTokenStr, _ := invalidClaimsToken.SignedString([]byte(validSecret))

	// invalid expired token
	expiredClaims := &claims{
		AuthUserID: "12345",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(-time.Second * 1).Unix(),
		},
	}
	expiredToken := jwt.NewWithClaims(jwt.SigningMethodHS256, expiredClaims)
	expiredTokenStr, _ := expiredToken.SignedString([]byte(validSecret))

	cases := map[string]struct {
		in         string
		authUserID string
		errmsg     string
	}{
		"valid-token":       {in: validTokenStr, authUserID: "12345"},
		"invalid-token":     {in: "invalid-token", errmsg: "token contains an invalid number of segments"},
		"invalid-claims":    {in: invalidClaimsTokenStr, errmsg: "Invalid token"},
		"invalid-expiresAt": {in: expiredTokenStr, errmsg: "token is expired by"},
	}

	for name, tt := range cases {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			claims, err := getClaims(tt.in)

			if tt.errmsg != "" {
				if !strings.Contains(err.Error(), tt.errmsg) {
					t.Errorf("expected error message to contain %q, but got %v", tt.errmsg, err.Error())
				}
			} else {
				authUserID := ""
				if claims != nil {
					authUserID = claims.AuthUserID
				}
				if authUserID != tt.authUserID {
					t.Errorf("want AuthUserID getClaims(%s) = %v, got %v", tt.in, tt.authUserID, claims.AuthUserID)
				}
			}
		})
	}
}
