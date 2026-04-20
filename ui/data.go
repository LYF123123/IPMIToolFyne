package ui

import (
	"fyne.io/fyne/v2"
)

type Item struct {
	Title, Intro string
	View         func(w fyne.Window) fyne.CanvasObject
}

var (
	Items = map[string]Item{
		"status": {Title: "Status", Intro: "System Status Summary", View: StatusScreen},
		"fans":   {Title: "Fans", Intro: "System Fans Control", View: FansScreen},
		"event":  {Title: "Event", Intro: "System Event Log", View: EventLogScreen},
	}

	ItemIndex = map[string][]string{
		"": {"status", "fans", "event"},
	}
)

// TODO
// new SDRs parameters list

var OnChangeFuncs []func()
