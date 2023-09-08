package handler

import (
	"database/sql"
	"encoding/json"
	"github.com/shshimamo/knowledge-auth/service"
	"net/http"
)

type AuthHandler interface {
	Signup(w http.ResponseWriter, r *http.Request)
	Signin(w http.ResponseWriter, r *http.Request)
	Signout(w http.ResponseWriter, r *http.Request)
}

type authHandler struct {
	db *sql.DB
}

func NewAuthHandler(db *sql.DB) AuthHandler {
	return &authHandler{db: db}
}

func (h *authHandler) Signup(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

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
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, "Unable to sign up", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *authHandler) Signin(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

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

	w.WriteHeader(http.StatusOK)
}

func (h *authHandler) Signout(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

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
