package api

import (
	"fmt"
	"manager/internal/data/datamodel"
	"manager/cmd/handlers"
	"manager/internal/publisher"
	"manager/utils"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type App struct {
	*mux.Router
	H *handlers.Handler
}

func NewApp() *App {
	dm, err := datamodel.New()
	if err != nil {
		utils.ErrLog.Falalf("%v", utils.DBConnError(err))
	}

	pub, err := publisher.NewDefaultPublisher()
	if err != nil {
		utils.ErrLog.Falalf("%v", utils.RedisConnErr(err))
	}

	app := &App{
		H:      handlers.New(dm, pub),
		Router: mux.NewRouter(),
	}

	app.register()
	return app
}

func (app App) NewServer() *http.Server {
	addr := fmt.Sprintf(":%s", utils.GetEnvVar("PORT"))
	utils.InfoLog.Printf("Serving flag configuration at %s", addr)
	
	return &http.Server{
		Addr:         addr,
		Handler:      app.Router,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
		IdleTimeout:  120 * time.Second,
		ErrorLog:     utils.ErrLog.Logger,
	}
}

func (app *App) register() {
	app.dashboardRoutes()
	app.providerRoutes()
	app.staticRoutes()
}
