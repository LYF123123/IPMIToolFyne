package main

import (
	"IPMITOOLFYNE/ui"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

const preferenceCurrentItem = "currentItem"

var topWindow fyne.Window

func main() {
	a := app.NewWithID("IPMI Tool")

	w := a.NewWindow("IPMI Tool")

	w.SetMainMenu(ui.MakeMenu(a, w))
	contentStage := container.NewStack()

	ui.ShowLoginDialog(a, w, contentStage)

	w.Resize(fyne.NewSize(800, 600))
	w.ShowAndRun()
}
