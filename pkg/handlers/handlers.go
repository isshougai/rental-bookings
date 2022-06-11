package handlers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/kugatsuno/udemy-modern-web-apps-golang/pkg/config"
	"github.com/kugatsuno/udemy-modern-web-apps-golang/pkg/models"
	"github.com/kugatsuno/udemy-modern-web-apps-golang/pkg/render"
)

// Repo is the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(c *gin.Context) {
	remoteIP := c.RemoteIP()
	session := sessions.Default(c)
	session.Set("remote_ip", remoteIP)
	session.Save()

	render.RenderTemplate(c, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(c *gin.Context) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"

	session := sessions.Default(c)
	var count int
	v := session.Get("count")
	if v == nil {
		count = 0
	} else {
		count = v.(int)
		count++
	}
	session.Set("count", count)
	session.Save()

	var remoteIP string
	ip := session.Get("remote_ip")
	if ip == nil {
		remoteIP = ""
	} else {
		remoteIP = ip.(string)
	}
	stringMap["remote_ip"] = remoteIP

	intMap := make(map[string]int)
	intMap["count"] = count

	render.RenderTemplate(c, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
		IntMap:    intMap,
	})
}
