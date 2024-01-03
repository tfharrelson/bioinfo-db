package main

import (
	"github.com/tfharrelson/bioinfo-db/internal/app"
)

func main() {

	app := app.App{}
	app.Initialize()

	app.Run()
}
