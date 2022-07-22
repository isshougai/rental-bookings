package render

import (
	"bytes"
	"fmt"
	"github.com/gin-contrib/sessions"
	"log"
	"path/filepath"
	"text/template"

	"github.com/gin-gonic/gin"
	"github.com/isshougai/rental-bookings/internal/config"
	"github.com/isshougai/rental-bookings/internal/models"
	"github.com/justinas/nosurf"
)

var functions = template.FuncMap{}

var app *config.AppConfig

// NewTemplaces sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData, c *gin.Context) *models.TemplateData {
	session := sessions.Default(c)
	td.Flash, _ = session.Get("flash").(string)
	td.Error, _ = session.Get("error").(string)
	td.Warning, _ = session.Get("warning").(string)
	td.CSRFToken = nosurf.Token(c.Request)
	session.Delete("flash")
	session.Delete("error")
	session.Delete("warning")
	err := session.Save()
	if err != nil {
		log.Println(err)
	}
	return td
}

// RenderTemplate renders templates using html templates
func RenderTemplate(c *gin.Context, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template

	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td, c)

	_ = t.Execute(buf, td)

	_, err := buf.WriteTo(c.Writer)
	if err != nil {
		fmt.Println("Error writing template to browser", err)
	}
}

// CreateTemplateCache creates a template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}
