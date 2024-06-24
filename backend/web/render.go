package main

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type templateData struct {
	Data map[string]any
}

func (app *application) render(w http.ResponseWriter, t string, td *templateData) {
	var tmpl *template.Template

	// if using cache then get template from map in the receiver
	if app.config.useCache {
		if templateFromMap, ok := app.templateMap[t]; ok {
			tmpl = templateFromMap
		}
	}

	if tmpl == nil {
		newTemplate, err := app.buildTemplateFromDisk(t)
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				// Handle 404 not found
				w.WriteHeader(http.StatusNotFound)
				app.render(w, "404.page.gohtml", nil)
				return
			}
			log.Println("Error building template: ", err)
		}
		log.Println("building template from disk.")
		tmpl = newTemplate
	}

	if td == nil {
		td = &templateData{}
	}

	if err := tmpl.ExecuteTemplate(w, t, td); err != nil {
		log.Println("Error executing template: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (app *application) buildTemplateFromDisk(t string) (*template.Template, error) {
	templateFiles := []string{
		"./templates/base.layout.gohtml",
		"./templates/partials/header.partial.gohtml",
		"./templates/partials/navbar.partial.gohtml",
		"./templates/partials/footer.partial.gohtml",
		fmt.Sprintf("./templates/%s", t),
	}

	tmpl, err := template.ParseFiles(templateFiles...)
	if err != nil {
		log.Printf("Error parsing templates: %v", err)
		return nil, err
	}

	app.templateMap[t] = tmpl

	return tmpl, nil
}
