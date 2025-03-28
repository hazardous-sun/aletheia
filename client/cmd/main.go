package main

import (
	"fact-ckert-client/gui"
	"fact-ckert-client/models"
	"fyne.io/fyne/v2/app"
)

func main() {
	config := models.NewConfig(false, true, true)
	appInstance := app.New()
	gui.Build(appInstance, config)
}
