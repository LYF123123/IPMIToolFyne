package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func MakeMenu(a fyne.App, w fyne.Window) *fyne.MainMenu {

	// help
	showAbout := func() {
		w := a.NewWindow("About")
		w.SetContent(widget.NewLabel("About IPMI Tool"))
		w.Show()
	}

	aboutItem := fyne.NewMenuItem("About", showAbout)

	helpMenu := fyne.NewMenu("Help", aboutItem)

	main := fyne.NewMainMenu(helpMenu)

	return main
}
