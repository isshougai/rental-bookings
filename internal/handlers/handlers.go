package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/isshougai/rental-bookings/internal/config"
	"github.com/isshougai/rental-bookings/internal/forms"
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
	err := session.Save()
	if err != nil {
		log.Println(err)
	}

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
	err := session.Save()
	if err != nil {
		log.Println(err)
	}

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
	var emptyReservation models.Reservation
	data := make(map[string]interface{})
	data["reservation"] = emptyReservation
	render.RenderTemplate(c, "make-reservation.page.tmpl", &models.TemplateData{
		Form: forms.New(nil),
		Data: data,
	})
}

// PostReservation handles the posting of a reservation form
func (m *Repository) PostReservation(c *gin.Context) {
	err := c.Request.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}

	reservation := models.Reservation{
		FirstName: c.PostForm("first_name"),
		LastName:  c.PostForm("last_name"),
		Email:     c.PostForm("email"),
		Phone:     c.PostForm("phone"),
	}

	form := forms.New(c.Request.PostForm)

	form.Required("first_name", "last_name", "email")
	form.MinLength("first_name", 3, c)
	form.IsEmail("email")

	if !form.Valid() {
		data := make(map[string]interface{})
		data["reservation"] = reservation

		render.RenderTemplate(c, "make-reservation.page.tmpl", &models.TemplateData{
			Form: form,
			Data: data,
		})

		return
	}

	session := sessions.Default(c)
	session.Set("reservation", reservation)
	err = session.Save()
	if err != nil {
		log.Println(err)
	}

	c.Redirect(http.StatusSeeOther, "/reservation-summary")
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
	_, err := c.Writer.WriteString(fmt.Sprintf("start date is %s and end date is %s", start, end))
	if err != nil {
		log.Println(err)
	}
}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

// AvailabilityJSON handles request for availability and send JSON response
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
	_, err = c.Writer.Write(out)
	if err != nil {
		log.Println(err)
	}
}

// Contact renders the Contact page
func (m *Repository) Contact(c *gin.Context) {
	render.RenderTemplate(c, "contact.page.tmpl", &models.TemplateData{})
}

// ReservationSummary returns the user the reservation summary
func (m *Repository) ReservationSummary(c *gin.Context) {
	session := sessions.Default(c)
	reservation, ok := session.Get("reservation").(models.Reservation)
	if !ok {
		log.Println("cannot get item from session")
		session.Set("error", "Can't get reservation from session")
		err := session.Save()
		if err != nil {
			log.Println(err)
		}
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}
	session.Delete("reservation")
	err := session.Save()
	if err != nil {
		log.Println(err)
	}

	data := make(map[string]interface{})
	data["reservation"] = reservation

	render.RenderTemplate(c, "reservation-summary.page.tmpl", &models.TemplateData{Data: data})
}
