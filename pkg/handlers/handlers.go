package handlers

import (
	"github.com/OusManDiouf/go-booking-app/pkg/config"
	"github.com/OusManDiouf/go-booking-app/pkg/models"
	"github.com/OusManDiouf/go-booking-app/pkg/render"
	"net/http"
)

var Repo *Repository

// Repository is the repository type
type Repository struct {
	AppConfig *config.AppConfig
}

// NewRepo create a new resository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		AppConfig: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

//----------------------------------------------------->

//Home handler
func (repo *Repository) Home(writer http.ResponseWriter, request *http.Request) {

	// Tsing session - Setting remoteIp adress
	remoteIp := request.RemoteAddr
	repo.AppConfig.SessionManager.Put(request.Context(), "remote_ip", remoteIp)

	render.Template(writer, "home.page.gohtml", &models.TemplateData{})
}

//About handler
func (repo *Repository) About(writer http.ResponseWriter, request *http.Request) {

	stringMap := make(map[string]string)
	stringMap["test"] = "Hello Lessou"

	// Pull out the session (TEST)
	remoteIp := repo.AppConfig.SessionManager.GetString(request.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIp

	render.Template(writer, "about.page.gohtml", &models.TemplateData{
		StringMap: stringMap,
	})
}
