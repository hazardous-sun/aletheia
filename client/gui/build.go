package gui

import (
	"fact-ckert-client/models"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"reflect"
)

func Build(a fyne.App, config *models.Config) {
	fields := collectConfiguredFields(*config)
	ctr := buildFields(fields)

	w := a.NewWindow("Client Test")
	w.SetContent(ctr)
	w.ShowAndRun()
}

func collectConfiguredFields(config models.Config) []string {
	v := reflect.ValueOf(config)
	typeOfS := v.Type()
	var result []string

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldName := typeOfS.Field(i).Name
		fieldType := field.Type()

		if fieldType.Kind() == reflect.Bool && field.Bool() {
			result = append(result, fieldName)
		}
	}

	return result
}

func buildFields(fields []string) fyne.CanvasObject {
	requiredWidgets := make([]fyne.CanvasObject, 0)
	requiredWidgets = buildRequiredFields(requiredWidgets)
	requiredCtr := container.NewGridWithRows(len(requiredWidgets), requiredWidgets...)

	optionalWidgets := make([]fyne.CanvasObject, 0)
	optionalWidgets = buildOptionalFields(optionalWidgets, fields)
	optionalCtr := container.NewGridWithRows(len(optionalWidgets), optionalWidgets...)

	return container.NewGridWithRows(2, requiredCtr, optionalCtr)
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
		case "Image":
			windowWidgets = append(windowWidgets, buildCheckField("Image:", flagsValues))
		case "Video":
			windowWidgets = append(windowWidgets, buildCheckField("Video:", flagsValues))
		case "Context":
			windowWidgets = append(windowWidgets, buildEntryContainerField("Context:"))
		default:
			fmt.Println("unknown")
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
