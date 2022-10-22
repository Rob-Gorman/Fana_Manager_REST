package api

import (
	"log"
	"manager/data/datamodel"
	"manager/handlers"
	"manager/publisher"
	"manager/utils"

	"github.com/gorilla/mux"
)

type App struct {
	*mux.Router
	H *handlers.Handler
}

func NewApp() *App {
	dm, err := datamodel.New()
	if err != nil {
		log.Fatal(utils.DBConnError(err))
	}

	pub, err := publisher.NewDefaultPublisher()
	if err != nil {
		log.Println(utils.RedisConnErr(err))
	}

	app := &App{
		H:      handlers.New(dm, pub),
		Router: mux.NewRouter(),
	}

	app.register()
	return app
}

func (app *App) register() {
	app.dashboardRoutes()
	app.providerRoutes()
	app.staticRoutes()
}
