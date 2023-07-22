package handler

import "github.com/shshimamo/knowledge-auth/model"

type SignupRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignupResponse struct {
	Token model.TokenStr `json:"token"`
}

type SigninRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SigninResponse struct {
	Token model.TokenStr `json:"token"`
}

type SignoutRequest struct {
	Token string `json:"token"`
}
