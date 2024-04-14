package handler

import (
	"fmt"
	"net/http"

	"github.com/SaiSoeSan/bookings/pkg/config"
	"github.com/SaiSoeSan/bookings/pkg/model"
	"github.com/SaiSoeSan/bookings/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository  {
	return &Repository{
		App : a,
	}
}

func NewHandler(r *Repository) {
	Repo = r
}

//Home is home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIp := r.RemoteAddr
	m.App.Session.Put(r.Context(),"remote_ip",remoteIp)
	
	render.RenderTempalate(w, "home.page.html",&model.TemplateData{})
}

//About is about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	//perfom logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello Again"

	remoteIp := m.App.Session.GetString(r.Context(),"remote_ip")
	fmt.Println(remoteIp)

	stringMap["remoteIp"] = remoteIp

	//send data to the template
	render.RenderTempalate(w,"about.page.html", &model.TemplateData{
		StringMap: stringMap,
	})
}