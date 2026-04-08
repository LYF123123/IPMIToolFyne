package ui

import "fyne.io/fyne/v2"

var OnChangeFuncs []func()

type Item struct {
	Title, Intro string
	View         func(w fyne.Window) fyne.CanvasObject
}

var (
	Items = map[string]Item{
		"System Status": {"System Status", "", systemStatusScreen},
	}
)
