package main

import (
	"net/http"

	"github.com/tfharrelson/bioinfo-db/internal/app"
)

func main() {
	router := app.CreateNewServer()

    http.ListenAndServe(":8080", router)
}
