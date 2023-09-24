package handler

import (
	"embed"
	"github.com/shshimamo/knowledge/backend/oauth/internal/util/bunapp"
	"net/http"
	"text/template"

	"github.com/uptrace/bunrouter"
)

//go:embed templates/*
var templates embed.FS

type WelcomeHandler struct {
	app *bunapp.App
	tpl *template.Template
}

func NewWelcomeHandler(app *bunapp.App) *WelcomeHandler {
	tpl, err := template.New("").ParseFS(templates, "templates/*.html")
	if err != nil {
		panic(err)
	}

	return &WelcomeHandler{
		app: app,
		tpl: tpl,
	}
}

func (h *WelcomeHandler) Welcome(w http.ResponseWriter, req bunrouter.Request) error {
	if err := h.tpl.ExecuteTemplate(w, "welcome.html", nil); err != nil {
		return err
	}
	return nil
}
