package model

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"strings"
	"testing"
	"time"
)

func setupJWTTokens() (validTokenStr, invalidTokenStr, invalidClaimsTokenStr, expiredTokenStr string) {
	validSecret := "secret" // TODO: ENV

	// valid token
	validClaims := &claims{
		AuthUserID: 12345,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 7).Unix(),
		},
	}
	validToken := jwt.NewWithClaims(jwt.SigningMethodHS256, validClaims)
	validTokenStr, _ = validToken.SignedString([]byte(validSecret))

	// invalid token
	invalidTokenStr = "invalid-token"

	// invalid claims
	invalidClaimsToken := jwt.New(jwt.SigningMethodHS256)
	invalidClaimsTokenStr, _ = invalidClaimsToken.SignedString([]byte(validSecret))

	// invalid expired token
	expiredClaims := &claims{
		AuthUserID: 12345,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(-time.Second * 1).Unix(),
		},
	}
	expiredToken := jwt.NewWithClaims(jwt.SigningMethodHS256, expiredClaims)
	expiredTokenStr, _ = expiredToken.SignedString([]byte(validSecret))

	return
}

func TestGetJWTToken(t *testing.T) {
	validTokenStr, invalidTokenStr, invalidClaimsTokenStr, expiredTokenStr := setupJWTTokens()

	tests := map[string]struct {
		in     string
		errmsg string
	}{
		"valid-token":       {in: validTokenStr},
		"invalid-token":     {in: invalidTokenStr, errmsg: "token contains an invalid number of segments aaa"},
		"invalid-claims":    {in: invalidClaimsTokenStr},
		"invalid-expiresAt": {in: expiredTokenStr, errmsg: "token is expired by"},
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got, err := getJWTToken(tt.in)

			if err != nil {
				if !strings.Contains(err.Error(), tt.errmsg) {
					t.Errorf("want: %q, got: %v", tt.errmsg, err.Error())
				}
			} else {
				if got == nil {
					t.Errorf("want: token, got: nil")
				}
			}
		})
	}
}

func TestGetClaims(t *testing.T) {
	validTokenStr, _, invalidClaimsTokenStr, _ := setupJWTTokens()
	validToken, _ := getJWTToken(validTokenStr)
	invalidClaimsToken, _ := getJWTToken(invalidClaimsTokenStr)

	tests := map[string]struct {
		in   *jwt.Token
		want *claims
	}{
		"valid-token":    {in: validToken, want: &claims{AuthUserID: 12345}},
		"invalid-claims": {in: invalidClaimsToken, want: nil},
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got, _ := getClaims(tt.in)

			if diff := cmp.Diff(got, tt.want, cmpopts.IgnoreTypes(jwt.StandardClaims{})); diff != "" {
				t.Errorf("want: %v, got: %v", tt.want, got)
			}
		})
	}
}
