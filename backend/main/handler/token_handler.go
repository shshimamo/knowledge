package handler

import (
	"encoding/json"
	"fmt"
	"github.com/shshimamo/knowledge/main/model"
	"log/slog"
	"net/http"
)

type TokenHandler interface {
	SetToken(w http.ResponseWriter, r *http.Request)
}

type tokenHandler struct {
	appEnv model.AppEnv
}

type SetTokenRequest struct {
	Token string `json:"token"`
}

func NewTokenHandler(appEnv model.AppEnv) TokenHandler {
	return &tokenHandler{appEnv: appEnv}
}

func (h *tokenHandler) SetToken(w http.ResponseWriter, r *http.Request) {
	slog.InfoContext(r.Context(), "SetToken called")

	if cookie, err := r.Cookie("token"); err == nil {
		tokenStr := cookie.Value
		fmt.Printf("\n%#v", tokenStr)
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var req SetTokenRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid parameter", http.StatusBadRequest)
		return
	}

	// MEMO: Only check if token is valid
	_, err = model.NewToken(req.Token)
	if err != nil {
		http.Error(w, "Invalid parameter", http.StatusBadRequest)
		return
	}

	//secure := true
	//if h.appEnv != model.Production {
	//	secure = false
	//}
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    string(req.Token),
		Path:     "/",
		HttpOnly: true,
		//Secure:   secure,
		Secure: false,
	})

	w.WriteHeader(http.StatusOK)
}
