package ui

import (
	"IPMITOOLFYNE/session"
	"fmt"
	"log"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func StatusScreen(_ fyne.Window) fyne.CanvasObject {

	sdrs := session.GetInstance().GetSDRs()
	log.Print("is sdrs ==nil? : ")
	log.Println(sdrs == nil)

	header := widget.NewLabelWithStyle("\n\nSystem Status", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	loading := container.NewCenter(container.NewVBox(widget.NewProgressBarInfinite(), widget.NewLabel("Data Syncing...")))
	list := widget.NewList(
		func() int { return len(sdrs) },
		func() fyne.CanvasObject {
			name := widget.NewLabelWithStyle("", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
			value := widget.NewLabelWithStyle("", fyne.TextAlignTrailing, fyne.TextStyle{Monospace: true})
			unit := widget.NewLabelWithStyle("", fyne.TextAlignLeading, fyne.TextStyle{Italic: true})
			return container.NewGridWithColumns(3, name, value, unit)
			// Create one container
		},
		func(id widget.ListItemID, o fyne.CanvasObject) {
			if id >= len(sdrs) {
				return
			}
			s := sdrs[id]

			container := o.(*fyne.Container)
			nameL := container.Objects[0].(*widget.Label)
			valueL := container.Objects[1].(*widget.Label)
			unitL := container.Objects[2].(*widget.Label)
			var valStr, unitStr string
			if s.Full != nil {
				valStr = fmt.Sprintf("%.3f", s.Full.SensorValue)
				unitStr = s.Full.SensorUnit.String()
			} else {
				valStr = "Discrete"
				unitStr = "Status"
			}
			nameL.SetText(s.SensorName())
			valueL.SetText(valStr)
			unitL.SetText(unitStr)
		},
	)
	listContainer := container.NewBorder(header, nil, nil, nil, list)
	statusStack := container.NewStack(loading, listContainer)
	go func() {
		ticker := time.NewTicker(100 * time.Millisecond)
		for range ticker.C {
			fyne.Do(func() {
				newData := session.GetInstance().GetSDRs()
				if newData != nil {
					sdrs = newData
					loading.Hide()
					list.Show()
					list.Refresh()
					statusStack.Refresh()
				}
			})
		}
	}()

	return statusStack
}
