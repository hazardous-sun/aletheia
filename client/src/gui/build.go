package gui

import (
	"aletheia-client/src/errors"
	"aletheia-client/src/models"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"strings"
)

var PostUrl string = ""
var Image bool = false
var Prompt string = ""
var Video bool = false

func Build(a fyne.App) {
	config, err := models.NewConfig()

	// Check if there were errors while generating the Config struct
	if err != nil {
		client_errors.Log(err.Error(), client_errors.ErrorLevel)
		return
	}

	ctr := buildFields(config)

	w := a.NewWindow("Client Test")
	w.SetContent(ctr)
	w.ShowAndRun()
}

func buildFields(config models.Config) fyne.CanvasObject {
	// Required widgets
	requiredWidgets := buildRequiredFields()

	// Optional widgets
	optionalWidgets := buildOptionalFields(config)

	widgetsCtr := container.NewGridWithRows(len(requiredWidgets)+len(optionalWidgets), append(requiredWidgets, optionalWidgets...)...)
	return container.NewBorder(
		widgetsCtr,
		nil, nil, nil,
		widget.NewButton("Send", func() {
			sendPackage(config)
		}),
	)
}

// Required fields
func buildRequiredFields() []fyne.CanvasObject {
	windowWidgets := make([]fyne.CanvasObject, 0)
	// Required fields
	urlField := buildEntryContainerField("URL:")
	windowWidgets = append(windowWidgets, urlField)
	return windowWidgets
}

// Optional fields
func buildOptionalFields(config models.Config) []fyne.CanvasObject {
	windowWidgets := make([]fyne.CanvasObject, 0)

	// Check if the Prompt entry field should be displayed
	if config.Prompt {
		windowWidgets = append(windowWidgets, buildEntryContainerField("Prompt:"))
	}

	// Check if the Image check field should be displayed
	if config.Image {
		windowWidgets = append(windowWidgets, buildCheckField(models.Image))
	}

	// Check if the Image check field should be displayed
	if config.Video {
		windowWidgets = append(windowWidgets, buildCheckField(models.Video))
	}

	return windowWidgets
}

// Constructors
func buildEntryContainerField(labelText string) fyne.CanvasObject {
	return container.NewBorder(
		nil, nil,
		widget.NewLabel(labelText), nil,
		widget.NewEntry(),
	)
}

func buildCheckField(labelText string) fyne.CanvasObject {
	var behavior func(bool)

	if labelText == models.Image {
		behavior = func(check bool) {
			Image = check
		}
	} else if labelText == models.Video {
		behavior = func(check bool) {
			Video = check
		}
	} else {
		panic("invalid field: " + labelText)
	}

	return widget.NewCheck(
		fmt.Sprintf("%s:", strings.ToTitle(labelText)),
		behavior,
	)
}

func sendPackage(config models.Config) {
	packageSent := models.PackageSent{
		Url:    PostUrl,
		Image:  Image,
		Prompt: Prompt,
		Video:  Video,
	}
	apiConnector := models.NewAPIConnector(PostUrl)
	response, err := apiConnector.SendPackage(fmt.Sprintf("localhost:%s", config.Port), packageSent)
	if err != nil {
		client_errors.Log(err.Error(), client_errors.ErrorLevel)
	} else {
		client_errors.Log(fmt.Sprintf("%v", response), client_errors.InfoLevel)
	}
}
