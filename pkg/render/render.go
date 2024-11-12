package render

import (
	"bytes"
	"github.com/giov27/bookings/pkg/config"
	"github.com/giov27/bookings/pkg/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

var app *config.AppConfig

// NewTemplates  set the configuration for the template pakcage
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(data *models.TemplateData) *models.TemplateData {

	return data
}

// RenderTemplate using render
func RenderTemplate(w http.ResponseWriter, tmpl string, data *models.TemplateData) {
	//create the template cache from the app config

	var tc map[string]*template.Template
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}
	//get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Error loading template:", tmpl)
	}

	buf := bytes.NewBuffer(nil)

	data = AddDefaultData(data)
	_ = t.Execute(buf, data)

	//render the template
	_, err := buf.WriteTo(w)
	if err != nil {
		log.Fatal("Error writing template:", err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	//	get all the files named *.page.tmpl from ./templates
	page, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return cache, err
	}

	// range all the files ending with *.page.tmpl
	for _, page := range page {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return cache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return cache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return cache, err
			}
		}

		cache[name] = ts

	}
	return cache, nil
}
