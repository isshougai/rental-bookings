package handlers

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/isshougai/rental-bookings/internal/config"
	"github.com/isshougai/rental-bookings/internal/models"
	"github.com/isshougai/rental-bookings/internal/render"
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

// Reservation renders the make a reservation page and displays form
func (m *Repository) Reservation(c *gin.Context) {
	render.RenderTemplate(c, "make-reservation.page.tmpl", &models.TemplateData{})
}

// Kiyomizu renders the Kiyomizu page
func (m *Repository) Kiyomizu(c *gin.Context) {
	render.RenderTemplate(c, "kiyomizu.page.tmpl", &models.TemplateData{})
}

// Gion renders the Gion page
func (m *Repository) Gion(c *gin.Context) {
	render.RenderTemplate(c, "gion.page.tmpl", &models.TemplateData{})
}

// Availability renders the Availability page
func (m *Repository) Availability(c *gin.Context) {
	render.RenderTemplate(c, "search-availability.page.tmpl", &models.TemplateData{})
}

// PostAvailability handles the post request for availability
func (m *Repository) PostAvailability(c *gin.Context) {
	start := c.PostForm("start")
	end := c.PostForm("end")
	c.Writer.WriteString(fmt.Sprintf("start date is %s and end date is %s", start, end))
}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

// AvailabilityJson handles request for availability and send JSON response
func (m *Repository) AvailabilityJSON(c *gin.Context) {
	resp := jsonResponse{
		OK:      true,
		Message: "Available!",
	}
	out, err := json.MarshalIndent(resp, "", "     ")
	if err != nil {
		log.Println(err)
	}

	c.Header("Content-Type", "application/json")
	c.Writer.Write(out)
}

// Contact renders the Contact page
func (m *Repository) Contact(c *gin.Context) {
	render.RenderTemplate(c, "contact.page.tmpl", &models.TemplateData{})
}
