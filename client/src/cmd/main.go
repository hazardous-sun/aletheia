package main

import (
	"aletheia-client/src/gui"
	"fyne.io/fyne/v2/app"
)

func main() {
	appInstance := app.New()
	gui.Build(appInstance)
}
