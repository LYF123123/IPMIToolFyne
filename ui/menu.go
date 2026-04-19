package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func MakeMenu(a fyne.App, w fyne.Window) *fyne.MainMenu {

	// help
	showAbout := func() {
		w := a.NewWindow("About")
		w.SetContent(container.NewVBox(
			widget.NewLabelWithStyle("IPMI Tool for Supermicro 4028", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
			widget.NewLabel("A simple utility for hardware monitoring and 4-zone fan control. Built with Go and Fyne."),
		))
		w.Show()
	}

	aboutItem := fyne.NewMenuItem("About", showAbout)

	helpMenu := fyne.NewMenu("Help", aboutItem)

	main := fyne.NewMainMenu(helpMenu)

	return main
}
