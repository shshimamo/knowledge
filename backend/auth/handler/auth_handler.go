package handler

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/shshimamo/knowledge-auth/model"
	"github.com/shshimamo/knowledge-auth/service"
	"net/http"
)

type AuthHandler interface {
	Signup(ctx context.Context, w http.ResponseWriter, r *http.Request)
	Signin(ctx context.Context, w http.ResponseWriter, r *http.Request)
	Signout(ctx context.Context, w http.ResponseWriter, r *http.Request)
}

type authHandler struct {
	db     *sql.DB
	appEnv model.AppEnv
}

func NewAuthHandler(db *sql.DB, appEnv model.AppEnv) AuthHandler {
	return &authHandler{db: db, appEnv: appEnv}
}

func (h *authHandler) Signup(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	var req SignupRequest
	// TODO: r.Body が空みたい
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Unable to sign up", http.StatusInternalServerError)
		return
	}

	u := service.NewAuthService(h.db)
	token, err := u.Signup(ctx, req.Email, req.Password)

	if err != nil {
		http.Error(w, "Unable to sign up", http.StatusInternalServerError)
		return
	}

	res := &SignupResponse{
		Token: token.Token,
	}
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, "Unable to sign up", http.StatusInternalServerError)
		return
	}

	secure := true
	if h.appEnv != model.Production {
		secure = false
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    string(token.Token),
		Path:     "/",
		HttpOnly: true,
		Secure:   secure,
	})
	w.WriteHeader(http.StatusOK)
}

func (h *authHandler) Signin(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	var req SigninRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Unable to sign in", http.StatusInternalServerError)
		return
	}

	u := service.NewAuthService(h.db)
	token, err := u.Signin(ctx, req.Email, req.Password)
	if err != nil {
		http.Error(w, "Unable to sign in", http.StatusInternalServerError)
		return
	}

	res := &SignupResponse{
		Token: token.Token,
	}
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, "Unable to sign up", http.StatusInternalServerError)
		return
	}

	secure := true
	if h.appEnv != model.Production {
		secure = false
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    string(token.Token),
		Path:     "/",
		HttpOnly: true,
		Secure:   secure,
	})
	w.WriteHeader(http.StatusOK)
}

func (h *authHandler) Signout(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	var req SignoutRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Unable to sign in", http.StatusInternalServerError)
		return
	}

	u := service.NewAuthService(h.db)
	err = u.Signout(ctx, req.Token)
	if err != nil {
		http.Error(w, "Unable to sign out", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
