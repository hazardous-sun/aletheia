package gui

import (
	"fact-ckert-client/src/models"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Build(a fyne.App, config *models.Config) {
	fields := models.CollectConfigValues(*config)
	ctr := buildFields(fields)

	w := a.NewWindow("Client Test")
	w.SetContent(ctr)
	w.ShowAndRun()
}

func buildFields(fields []string) fyne.CanvasObject {
	// Required widgets
	requiredWidgets := make([]fyne.CanvasObject, 0)
	requiredWidgets = buildRequiredFields(requiredWidgets)

	// Optional widgets
	optionalWidgets := make([]fyne.CanvasObject, 0)
	optionalWidgets = buildOptionalFields(optionalWidgets, fields)

	widgetsCtr := container.NewGridWithRows(len(requiredWidgets)+len(optionalWidgets), append(requiredWidgets, optionalWidgets...)...)
	return container.NewBorder(
		widgetsCtr,
		nil, nil, nil,
		widget.NewButton("Send", func() {
			fmt.Println(fmt.Sprintf("Sending"))
		}),
	)
}

// Required fields
func buildRequiredFields(windowWidgets []fyne.CanvasObject) []fyne.CanvasObject {
	// Required fields
	urlField := buildEntryContainerField("URL:")
	windowWidgets = append(windowWidgets, urlField)
	return windowWidgets
}

// Optional fields
func buildOptionalFields(windowWidgets []fyne.CanvasObject, fields []string) []fyne.CanvasObject {
	flagsValues := models.NewConfig(false, false, false)
	for _, v := range fields {
		switch v {
		case "CONTEXT":
			windowWidgets = append(windowWidgets, buildEntryContainerField("Context:"))
		case "IMAGE":
			windowWidgets = append(windowWidgets, buildCheckField("Image:", flagsValues))
		case "VIDEO":
			windowWidgets = append(windowWidgets, buildCheckField("Video:", flagsValues))
		default:
			continue
		}
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

func buildCheckField(labelText string, flagsValues *models.Config) fyne.CanvasObject {
	return widget.NewCheck(
		labelText,
		func(checked bool) {
			flagsValues.Image = checked
		},
	)
}
