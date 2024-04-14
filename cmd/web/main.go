package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/SaiSoeSan/bookings/pkg/config"
	"github.com/SaiSoeSan/bookings/pkg/handler"
	"github.com/SaiSoeSan/bookings/pkg/render"
	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"
var app config.AppConfig
var session *scs.SessionManager

func main() {

	//change this to true in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	//put into config
	app.Session = session

	tc, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handler.NewRepo(&app)

	render.NewTemplates(&app)
	handler.NewHandler(repo)

	fmt.Println("port Number", portNumber)


	server := &http.Server {
		Addr : portNumber,
		Handler : routes(&app),
	}

	err = server.ListenAndServe()
	log.Fatal(err)
}