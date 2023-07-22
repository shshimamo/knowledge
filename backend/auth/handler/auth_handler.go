package handler

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/shshimamo/knowledge-auth/service"
	"net/http"
)

type AuthHandler interface {
	Signup(ctx context.Context, w http.ResponseWriter, r *http.Request)
	Signin(ctx context.Context, w http.ResponseWriter, r *http.Request)
	Signout(ctx context.Context, w http.ResponseWriter, r *http.Request)
}

type authHandler struct {
	db *sql.DB
}

func New(db *sql.DB) AuthHandler {
	return &authHandler{db: db}
}

func (h *authHandler) Signup(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	var req SignupRequest
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
	json.NewEncoder(w).Encode(res)
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

	res := &SigninResponse{
		Token: token.Token,
	}
	json.NewEncoder(w).Encode(res)
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
