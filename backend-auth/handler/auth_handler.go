package handler

import (
	"database/sql"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/shshimamo/knowledge-auth/model"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type AuthHandler interface {
	Signup(w http.ResponseWriter, r *http.Request)
	Signin(w http.ResponseWriter, r *http.Request)
	Signout(w http.ResponseWriter, r *http.Request)
}

type authHandler struct {
	db *sql.DB
}

func New(db *sql.DB) AuthHandler {
	return &authHandler{db: db}
}

type Claims struct {
	AuthUserID int `json:"auth_user_id"`
	jwt.StandardClaims
}

func (h *authHandler) Signup(w http.ResponseWriter, r *http.Request) {
	var authUser model.AuthUser
	_ = json.NewDecoder(r.Body).Decode(&authUser)

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(authUser.Password), 8)

	tx, err := h.db.Begin()
	if err != nil {
		http.Error(w, "Unable to start transaction", http.StatusInternalServerError)
		return
	}
	defer tx.Rollback()

	err = tx.QueryRow("INSERT INTO auth_users (email, password_digest) VALUES ($1, $2) RETURNING id", authUser.Email, hashedPassword).Scan(&authUser.ID)
	if err != nil {
		http.Error(w, "Unable to sign up", http.StatusBadRequest)
		return
	}

	token := jwt.New(jwt.SigningMethodHS256)
	expiresAt := time.Now().Add(time.Hour * 24 * 7)
	claims := &Claims{
		AuthUserID: authUser.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt.Unix(),
		},
	}

	token.Claims = claims
	tokenString, _ := token.SignedString([]byte("secret")) // TODO: 環境変数にする

	_, err = tx.Query("INSERT INTO tokens (auth_user_id, token, active, expires_at) VALUES ($1, $2, $3, $4)", authUser.ID, tokenString, true, expiresAt)
	if err != nil {
		http.Error(w, "Unable to create token", http.StatusBadRequest)
		return
	}

	err = tx.Commit()
	if err != nil {
		http.Error(w, "Unable to commit transaction", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(tokenString)
}

func (h *authHandler) Signin(w http.ResponseWriter, r *http.Request) {
	var authUser model.AuthUser
	_ = json.NewDecoder(r.Body).Decode(&authUser)

	result := h.db.QueryRow("SELECT id, password_digest FROM auth_users WHERE email=$1", authUser.Email)
	err := result.Scan(&authUser.ID, &authUser.PasswordDigest)
	if err != nil {
		http.Error(w, "AuthUser not found", http.StatusBadRequest)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(authUser.PasswordDigest), []byte(authUser.Password))
	if err != nil {
		http.Error(w, "Invalid password", http.StatusBadRequest)
		return
	}

	token := jwt.New(jwt.SigningMethodHS256)
	expiresAt := time.Now().Add(time.Hour * 24 * 7)
	claims := &Claims{
		AuthUserID: authUser.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt.Unix(),
		},
	}

	token.Claims = claims
	tokenString, _ := token.SignedString([]byte("secret"))

	_, err = h.db.Query("INSERT INTO tokens (auth_user_id, token, active, expires_at) VALUES ($1, $2, $3, $4)", authUser.ID, tokenString, true, expiresAt)
	if err != nil {
		http.Error(w, "Unable to create token", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(tokenString)
}

func (h *authHandler) Signout(w http.ResponseWriter, r *http.Request) {
	var token model.Token
	_ = json.NewDecoder(r.Body).Decode(&token)

	_, err := h.db.Query("UPDATE tokens SET active=$1 WHERE token=$2", false, token.Token)
	if err != nil {
		http.Error(w, "Unable to sign out", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
