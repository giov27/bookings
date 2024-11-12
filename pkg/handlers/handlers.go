package handlers

import (
	"fmt"
	"github.com/giov27/bookings/pkg/config"
	"github.com/giov27/bookings/pkg/models"
	"github.com/giov27/bookings/pkg/render"
	"net/http"
)

// Repo the repository used by the handler
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepository creates a new repository
func NewRepository(app *config.AppConfig) *Repository {
	return &Repository{App: app}
}

// NewHandler sets the repository for the handler
func NewHandler(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIp := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIp)
	fmt.Println("from home", remoteIp)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Holla from handlers yuhu!"

	remoteIp := m.App.Session.GetString(r.Context(), "remote_ip")
	fmt.Println(remoteIp)
	stringMap["remote_ip"] = remoteIp

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
