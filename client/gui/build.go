package gui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Build(app fyne.App) {
	w := app.NewWindow("GUI")
	w.Resize(fyne.NewSize(800, 600))

	displayCtr := container.NewBorder(
		nil, nil, nil, nil,
		widget.NewButton(
			"CLICK",
			func() {
				fmt.Println("clicked")
			},
		),
	)

	w.SetContent(displayCtr)

	w.ShowAndRun()
}
