package main

import (
	"fyne.io/fyne/v2/app"
	"mobile-client/gui"
)

func main() {
	app := app.New()
	gui.Build(app)
}
