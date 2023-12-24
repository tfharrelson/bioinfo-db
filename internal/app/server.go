package app

import (
	"github.com/gorilla/mux"
)

func CreateNewServer() *mux.Router {
    // this is an exported function that the main app will use to create a new server
    // Create the router
    router := mux.NewRouter()

    // create the routes using our handy dandy helper function
    CreateRoutes(router)

    // return the router
    return router
}


