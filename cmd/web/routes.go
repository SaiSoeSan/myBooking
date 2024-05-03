package main

import (
	"net/http"

	"github.com/SaiSoeSan/bookings/pkg/config"
	"github.com/SaiSoeSan/bookings/pkg/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {

	mux := chi.NewRouter()

	//middlewares
	mux.Use(middleware.Recoverer)
	mux.Use(WriteToConsole)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)
	
	mux.Get("/", handler.Repo.Home)
	mux.Get("/about",handler.Repo.About)

	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/*",http.StripPrefix("/static",fileServer))

	return mux
}
