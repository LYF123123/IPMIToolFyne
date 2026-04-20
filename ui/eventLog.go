package ui

import (
	"IPMITOOLFYNE/session"
	"strconv"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func EventLogScreen(win fyne.Window) fyne.CanvasObject {

	allSEL := session.GetInstance().GetSELs()

	title := widget.NewLabelWithStyle("BMC SENSOR MONITOR", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	title.Importance = widget.HighImportance

	columnTitle := container.NewGridWithColumns(4,
		widget.NewLabelWithStyle("Id", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		widget.NewLabelWithStyle("Time", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		widget.NewLabelWithStyle("Level", fyne.TextAlignTrailing, fyne.TextStyle{Bold: true}),
		widget.NewLabelWithStyle("Event", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
	)

	topContainer := container.NewVBox(
		container.NewPadded(title),
		widget.NewSeparator(),
		container.NewPadded(columnTitle),
	)

	list := widget.NewList(
		func() int {
			return len(allSEL)
		},
		func() fyne.CanvasObject {
			idL := widget.NewLabelWithStyle("", fyne.TextAlignLeading, fyne.TextStyle{Monospace: true})
			timeL := widget.NewLabelWithStyle("", fyne.TextAlignLeading, fyne.TextStyle{Monospace: true})
			levelL := widget.NewLabelWithStyle("", fyne.TextAlignTrailing, fyne.TextStyle{Bold: true})
			eventL := widget.NewLabel("")
			return container.NewGridWithColumns(4, idL, timeL, levelL, eventL)
		},
		func(id widget.ListItemID, o fyne.CanvasObject) {
			idx := len(allSEL) - 1 - int(id)
			if idx < 0 {
				return
			}
			item := allSEL[id]
			container := o.(*fyne.Container)
			idL := container.Objects[0].(*widget.Label)
			timeL := container.Objects[1].(*widget.Label)
			levelL := container.Objects[2].(*widget.Label)
			eventL := container.Objects[3].(*widget.Label)

			idL.Importance = widget.MediumImportance
			timeL.Importance = widget.MediumImportance
			levelL.Importance = widget.MediumImportance
			eventL.Importance = widget.MediumImportance

			idL.SetText(strconv.Itoa(id+1))
			severityStr := "Unknown"
			eventDesc := "OEM/Unknown Record"
			timeStr := "Unknown"
			if item.Standard != nil {
				severityStr = string(item.Standard.EventSeverity())
				eventDesc = item.Standard.EventString()
				timeStr = item.Standard.Timestamp.String()
			} else if item.OEMNonTimestamped != nil {
				severityStr = "OEM"
				eventDesc = "Custom OEM Data"
				timeStr = item.OEMTimestamped.Timestamp.String()
			}
			timeL.SetText(timeStr)
			levelL.SetText(severityStr)
			if strings.Contains(strings.ToLower(string(severityStr)), "critical") {
				idL.Importance = widget.DangerImportance
				timeL.Importance = widget.DangerImportance
				levelL.Importance = widget.DangerImportance
				eventL.Importance = widget.DangerImportance
			} else if strings.Contains(strings.ToLower(string(severityStr)), "warning") {
				idL.Importance = widget.WarningImportance
				timeL.Importance = widget.WarningImportance
				levelL.Importance = widget.WarningImportance
				eventL.Importance = widget.WarningImportance
			}
			eventL.SetText(eventDesc)
		},
	)
	loading := container.NewCenter(container.NewVBox(widget.NewProgressBarInfinite(), widget.NewLabel("Data Syncing...")))
	listContainer := container.NewBorder(topContainer, nil, nil, nil, list)
	statusStack := container.NewStack(loading, listContainer)
	go func() {
		ticker := time.NewTicker(500 * time.Millisecond)
		for range ticker.C {
			fyne.Do(func() {
				newAllSEL := session.GetInstance().GetSELs()
				if len(newAllSEL) > 0 {
					allSEL = newAllSEL
					loading.Hide()
					list.Show()
					list.Refresh()
				}
			})
		}
	}()

	return statusStack
}
