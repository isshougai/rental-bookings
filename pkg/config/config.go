package config

import (
	"log"
	"text/template"

	"github.com/gin-contrib/sessions/cookie"
)

// AppConfig holds the application config
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	InProduction  bool
	Store         cookie.Store
}
