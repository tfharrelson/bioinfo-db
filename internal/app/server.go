package app

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	_ "github.com/lib/pq"
)

type App struct {
	Router *mux.Router
	db *sql.DB
}

func (a *App) Initialize() {
	// this is an exported function that the main app will use to create a new server
	// Create the router
	router := mux.NewRouter()

	// create the routes using our handy dandy helper function
	CreateRoutes(router)

	// return the router
	a.Router = router

	// start up the DB connection
	// TODO: break apart this bad boi and do some error checking to make sure the env vars are set
	connString := "postgres://" + os.Getenv("DBUSER") + ":" + os.Getenv("DBPASSWORD") + "@localhost/" + os.Getenv("DBNAME") + "?sslmode=disable"
	db, err := sql.Open(connString, "postgres")
	if err != nil {
		panic("failed db connection")
	}
	a.db = db
}

func (a *App) Run() {
	log.Fatal(http.ListenAndServe(":8080", a.Router))
}
