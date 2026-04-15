package ui

import (
	"IPMITOOLFYNE/session"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func InitMainUI(a fyne.App, w fyne.Window, stage *fyne.Container) {
	w.SetMaster()
	// Light and Dark change
	themes := container.NewGridWithColumns(2,
		widget.NewButton("Dark", func() {
			a.Settings().SetTheme(&ForcedVariant{
				Theme:   theme.DefaultTheme(),
				Variant: theme.VariantDark,
			})
		}),
		widget.NewButton("Light", func() {
			a.Settings().SetTheme(&ForcedVariant{
				Theme:   theme.DefaultTheme(),
				Variant: theme.VariantLight,
			})
		}),
	)
	// left Nav
	ids := ItemIndex[""]
	navList := widget.NewList(
		func() int {
			return len(ids)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("Menu Template")
		},
		func(id widget.ListItemID, o fyne.CanvasObject) {
			key := ids[id]
			o.(*widget.Label).SetText(Items[key].Title)
		},
	)
	navList.OnSelected = func(id widget.ListItemID) {
		key := ids[id]
		item := Items[key]

		for _, f := range OnChangeFuncs {
			f()
		}
		OnChangeFuncs = nil
		stage.Objects = []fyne.CanvasObject{item.View(w)}
		stage.Refresh()
	}
	leftNav := container.NewBorder(nil, themes, nil, nil, navList)
	split := container.NewHSplit(leftNav, stage)
	split.Offset = 0.2
	w.SetContent(split)
	navList.Select(0)
	// Set close hook
	w.SetCloseIntercept(func() {
		for _, f := range OnChangeFuncs {
			f()
		}
		dialog.ShowConfirm("Exit", "Exit and Close connection?", func(ok bool) {
			if ok {
				session.GetInstance().Logout()
				w.Close()
			}
		}, w)
	})
}
