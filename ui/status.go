package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func StatusScreen(_ fyne.Window) fyne.CanvasObject {
	content := container.NewVBox(
		widget.NewLabelWithStyle("\n\nWelcome to the Fyne toolkit demo app", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		widget.NewLabelWithStyle("\nWith great thanks to our many kind sponsors\n", fyne.TextAlignCenter, fyne.TextStyle{Italic: true}))
	return content
}
