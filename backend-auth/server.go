package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

type AuthUser struct {
	ID             int    `json:"-"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	PasswordDigest string `json:"-"`
}

type Token struct {
	ID         int
	AuthUserID int
	Token      string
	Active     bool
}

type Claims struct {
	AuthUserID int `json:"auth_user_id"`
	jwt.StandardClaims
}

type databaseConfig struct {
	host     string
	port     string
	user     string
	password string
	dbname   string
}

type AppEnv string

const (
	Production AppEnv = "production"
)

var db *sql.DB

func main() {
	var err error

	appEnv := AppEnv(os.Getenv("APP_ENV"))

	db, err = setupDatabase(appEnv)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	http.HandleFunc("/signup", signup)
	http.HandleFunc("/signin", signin)
	http.HandleFunc("/signout", signout)

	log.Fatal(http.ListenAndServe(":80", nil))
}

func signup(w http.ResponseWriter, r *http.Request) {
	var authUser AuthUser
	_ = json.NewDecoder(r.Body).Decode(&authUser)

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(authUser.Password), 8)

	tx, err := db.Begin()
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

func signin(w http.ResponseWriter, r *http.Request) {
	var authUser AuthUser
	_ = json.NewDecoder(r.Body).Decode(&authUser)

	result := db.QueryRow("SELECT id, password_digest FROM auth_users WHERE email=$1", authUser.Email)
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

	_, err = db.Query("INSERT INTO tokens (auth_user_id, token, active, expires_at) VALUES ($1, $2, $3, $4)", authUser.ID, tokenString, true, expiresAt)
	if err != nil {
		http.Error(w, "Unable to create token", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(tokenString)
}

func signout(w http.ResponseWriter, r *http.Request) {
	var token Token
	_ = json.NewDecoder(r.Body).Decode(&token)

	_, err := db.Query("UPDATE tokens SET active=$1 WHERE token=$2", false, token.Token)
	if err != nil {
		http.Error(w, "Unable to sign out", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func setupDatabase(env AppEnv) (*sql.DB, error) {
	var dbCfg databaseConfig
	if env == Production {
		dbCfg = databaseConfig{
			host:     os.Getenv("DB_HOST"),
			port:     os.Getenv("DB_PORT"),
			user:     os.Getenv("DB_USER"),
			password: os.Getenv("DB_PASSWORD"),
			dbname:   os.Getenv("DB_NAME"),
		}
	} else {
		dbCfg = databaseConfig{
			host:     "localhost",
			port:     "5432",
			user:     "postgres",
			password: "password",
			dbname:   "knowledge",
		}
	}
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbCfg.host, dbCfg.port, dbCfg.user, dbCfg.password, dbCfg.dbname)
	db, err := sql.Open("postgres", connStr)

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, err
}
