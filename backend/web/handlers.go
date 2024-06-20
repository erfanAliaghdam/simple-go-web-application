package main

import (
	"net/http"
)

func (app *application) Index(
	w http.ResponseWriter,
	r *http.Request) {
	app.render(w, "index.page.gohtml", nil)
}
