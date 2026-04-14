package main

import (
	"IPMITOOLFYNE/ui"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

const preferenceCurrentItem = "currentItem"

var topWindow fyne.Window

func main() {
	a := app.NewWithID("IPMI Tool")

	w := a.NewWindow("IPMI Tool")

	w.SetMainMenu(ui.MakeMenu(a, w))
	w.SetMaster()

	contentStage := container.NewStack()

	// Light and Dark change
	themes := container.NewGridWithColumns(2,
		widget.NewButton("Dark", func() {
			a.Settings().SetTheme(&ui.ForcedVariant{
				Theme:   theme.DefaultTheme(),
				Variant: theme.VariantDark,
			})
		}),
		widget.NewButton("Light", func() {
			a.Settings().SetTheme(&ui.ForcedVariant{
				Theme:   theme.DefaultTheme(),
				Variant: theme.VariantLight,
			})
		}),
	)
	// left Nav
	ids := ui.ItemIndex[""]
	navList := widget.NewList(
		func() int {
			return len(ids)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("Menu Template")
		},
		func(id widget.ListItemID, o fyne.CanvasObject) {
			key := ids[id]
			o.(*widget.Label).SetText(ui.Items[key].Title)
		},
	)
	navList.OnSelected = func(id widget.ListItemID) {
		key := ids[id]
		item := ui.Items[key]

		for _, f := range ui.OnChangeFuncs {
			f()
		}
		ui.OnChangeFuncs = nil
		contentStage.Objects = []fyne.CanvasObject{item.View(w)}
		contentStage.Refresh()
	}
	leftNav := container.NewBorder(nil, themes, nil, nil, navList)
	split := container.NewHSplit(leftNav, contentStage)
	split.Offset = 0.2
	w.SetContent(split)
	navList.Select(0)

	w.Resize(fyne.NewSize(800, 600))
	w.ShowAndRun()
}
