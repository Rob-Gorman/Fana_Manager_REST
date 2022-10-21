package api

import (
	"manager/database"
	"manager/handlers"

	"github.com/gorilla/mux"
)

type App struct {
	*mux.Router // this is our express-router
	H           handlers.Handler
	// this Handler is a wrapper around our DB to allow us to define methods on it
	// those methods being our controller functions (handlers)
}

func NewApp() *App {
	s := &App{
		H:      handlers.New(database.Init()),
		Router: mux.NewRouter(),
	}

	s.routes()
	return s
}

func (s *App) routes() {
	s.dashboardRoutes()
	s.providerRoutes()
	s.staticRoutes()
}
