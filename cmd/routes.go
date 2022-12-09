package main

import (
	"github.com/Cuthbert20/go-trivia/handlers"
)

func setupRoutes(app *filber.App) {

	app.Get("/", handlers.Home)

}
