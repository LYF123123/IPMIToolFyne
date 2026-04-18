package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func FansScreen(win fyne.Window) fyne.CanvasObject {
	content := container.NewVBox(
		widget.NewLabelWithStyle("\n\nFans Status", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		widget.NewLabelWithStyle("\n\n", fyne.TextAlignCenter, fyne.TextStyle{Italic: true}))
	return content
}
