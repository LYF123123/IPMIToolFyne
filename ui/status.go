package ui

import (
	"IPMITOOLFYNE/session"
	"fmt"
	"log"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func StatusScreen(_ fyne.Window) fyne.CanvasObject {

	sdrs := session.GetInstance().GetSDRs()
	log.Print("is sdrs == nil? : ")
	log.Println(sdrs == nil)

	title := widget.NewLabelWithStyle("BMC SENSOR MONITOR", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	title.Importance = widget.HighImportance

	columnTitle := container.NewGridWithColumns(4,
		widget.NewLabelWithStyle("Sensor Name", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		widget.NewLabelWithStyle("Reading", fyne.TextAlignTrailing, fyne.TextStyle{Bold: true}),
		widget.NewLabelWithStyle("Unit", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		widget.NewLabelWithStyle("Status", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
	)

	topContainer := container.NewVBox(
		container.NewPadded(title),
		widget.NewSeparator(),
		container.NewPadded(columnTitle),
	)

	loading := container.NewCenter(container.NewVBox(widget.NewProgressBarInfinite(), widget.NewLabel("Data Syncing...")))
	list := widget.NewList(
		func() int { return len(sdrs) },
		func() fyne.CanvasObject {
			name := widget.NewLabelWithStyle("", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
			value := widget.NewLabelWithStyle("", fyne.TextAlignTrailing, fyne.TextStyle{Monospace: true})
			unit := widget.NewLabelWithStyle("", fyne.TextAlignLeading, fyne.TextStyle{Italic: true})
			status := widget.NewLabelWithStyle("", fyne.TextAlignLeading, fyne.TextStyle{Italic: true})
			return container.NewGridWithColumns(4, name, value, unit, status)
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
			statusL := container.Objects[3].(*widget.Label)
			// Reset Importance status
			nameL.Importance = widget.MediumImportance
			valueL.Importance = widget.MediumImportance
			unitL.Importance = widget.MediumImportance
			statusL.Importance = widget.MediumImportance

			var valStr, unitStr, statusStr string
			if s.Full != nil {
				valStr = fmt.Sprintf("%.3f", s.Full.SensorValue)
				unitStr = s.Full.SensorUnit.String()
				statusStr = s.Full.SensorStatus
				if strings.Contains(s.SensorName(), "Intru") && s.Full.SensorValue != 0 {
					// Chassis Intru
					nameL.Importance = widget.DangerImportance
					valueL.Importance = widget.DangerImportance
					unitL.Importance = widget.DangerImportance
					statusL.Importance = widget.DangerImportance
				}
			} else {
				valStr = "Discrete"
				unitStr = "       "
				statusStr = "Unknown"
			}

			if statusStr == "Unknown" {
				// What happen to this SDR????
				nameL.Importance = widget.WarningImportance
				valueL.Importance = widget.WarningImportance
				unitL.Importance = widget.WarningImportance
				statusL.Importance = widget.WarningImportance
			}
			nameL.SetText(s.SensorName())
			valueL.SetText(valStr)
			unitL.SetText(unitStr)
			statusL.SetText(statusStr)
		},
	)
	listContainer := container.NewBorder(topContainer, nil, nil, nil, list)
	statusStack := container.NewStack(loading, listContainer)
	go func() {
		ticker := time.NewTicker(500 * time.Millisecond)
		for range ticker.C {
			fyne.Do(func() {
				newData := session.GetInstance().GetSDRs()
				if len(newData) > 0 {
					sdrs = newData
					loading.Hide()
					list.Show()
					list.Refresh()
					// statusStack.Refresh()
				}
			})
		}
	}()

	return statusStack
}
