package gui

import (
	"aletheia-client/src/errors"
	"aletheia-client/src/models"
	"bytes"
	"encoding/json"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"io"
	"net/http"
	"strings"
)

var PostUrl string = ""
var Image bool = false
var Video bool = false
var Prompt string = ""
var promptEntry *widget.Entry
var answerBox *widget.Entry

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
	requiredWidgets := buildRequiredFields()
	optionalWidgets := buildOptionalFields(config)

	answerBox = widget.NewMultiLineEntry()
	answerBox.SetPlaceHolder("Waiting for server response...")
	answerBox.Hide()

	allWidgets := append(requiredWidgets, optionalWidgets...)
	allWidgets = append(allWidgets, answerBox)

	formContainer := container.NewGridWithRows(len(allWidgets), allWidgets...)

	return container.NewBorder(
		formContainer,
		nil, nil, nil,
		widget.NewButton("Send", func() {
			Prompt = promptEntry.Text
			sendPackage(config)
		}),
	)
}

// Required fields
func buildRequiredFields() []fyne.CanvasObject {
	promptEntry = widget.NewMultiLineEntry()
	promptEntry.SetPlaceHolder("Enter your search query...")
	return []fyne.CanvasObject{
		widget.NewLabel("Prompt"),
		promptEntry,
	}
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
	requestBody := map[string]interface{}{
		"query":        Prompt,
		"pagesToVisit": 5,
	}

	bodyJson, err := json.Marshal(requestBody)
	if err != nil {
		client_errors.Log("Failed to encode request body: "+err.Error(), client_errors.ErrorLevel)
		answerBox.SetText("Error preparing request.")
		answerBox.Show()
		return
	}

	// Log the package being sent
	client_errors.Log("Sending JSON to server:\n"+string(bodyJson), client_errors.InfoLevel)

	apiURL := "http://localhost:" + config.Port + "/crawl"
	resp, err := http.Post(apiURL, "application/json", bytes.NewReader(bodyJson))
	if err != nil {
		client_errors.Log("Request failed: "+err.Error(), client_errors.ErrorLevel)
		answerBox.SetText("Error: " + err.Error())
		answerBox.Show()
		return
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		client_errors.Log("Reading response failed: "+err.Error(), client_errors.ErrorLevel)
		answerBox.SetText("Error reading response.")
		answerBox.Show()
		return
	}

	client_errors.Log("Received response: "+string(respBody), client_errors.InfoLevel)
	answerBox.SetText(string(respBody))
	answerBox.Show()
}
