package app

import (
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
}

func (a *App) Initialize() {
	// this is an exported function that the main app will use to create a new server
	// Create the router
	router := mux.NewRouter()

	// create the routes using our handy dandy helper function
	CreateRoutes(router)

	// return the router
	a.Router = router
}

func (a *App) Run() {
	http.ListenAndServe(":8080", a.Router)
}
