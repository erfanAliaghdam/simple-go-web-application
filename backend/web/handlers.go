package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (app *application) Index(
	w http.ResponseWriter,
	r *http.Request) {
	app.render(w, "index.page.gohtml", nil)
}

func (app *application) ShowPage(
	w http.ResponseWriter,
	r *http.Request) {
	pageName := chi.URLParam(r, "page")
	app.render(w, fmt.Sprintf("%s.page.gohtml", pageName), nil)
}
