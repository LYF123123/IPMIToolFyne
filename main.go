package main

import (
	"IPMITOOLFYNE/ui"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var topWindow fyne.Window

func main() {
	a := app.NewWithID("IPMI Tool")

	w := a.NewWindow("IPMI Tool")

	w.SetMainMenu(ui.MakeMenu(a, w))
	w.SetMaster()

	content := container.NewStack()
	title := widget.NewLabel("Component name")
	intro := widget.NewLabel("An introduction would probably go\nhere, as well as a")
	intro.Wrapping = fyne.TextWrapWord

	top := container.NewVBox(title, widget.NewSeparator(), intro)
	setItem := func(i ui.Item) {
		if fyne.CurrentDevice().IsMobile() {
			child := a.NewWindow(i.Title)
			topWindow = child
			child.SetContent(i.View(topWindow))
			child.Show()
			child.SetOnClosed(func() {
				topWindow = w
			})
			return
		}
		title.SetText(i.Title)
		isMarkdown := len(i.Intro) == 0
		if isMarkdown {
			intro.SetText(i.Intro)
		}

		if i.Title == "System Status" || isMarkdown {
			top.Hide()
		} else {
			top.Show()
		}

		content.Objects = []fyne.CanvasObject{i.View(w)}
		content.Refresh()
	}
	item:=container.NewBorder(top,nil,nil,nil,content)
	if fyne.CurrentDevice().IsMobile(){
		w.SetContent(make)
	}





	w.Resize(fyne.NewSize(640, 460))
	w.ShowAndRun()
}
