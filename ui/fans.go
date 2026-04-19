package ui

import (
	"IPMITOOLFYNE/session"
	"fmt"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/bougou/go-ipmi"
)

func FansScreen(win fyne.Window) fyne.CanvasObject {

	allSDRs := session.GetInstance().GetSDRs()

	fanSDRs := getFanSDRs(allSDRs)

	fanRPMList := widget.NewList(
		func() int { return len(fanSDRs) },
		func() fyne.CanvasObject {
			name := widget.NewLabelWithStyle("", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
			value := widget.NewLabelWithStyle("", fyne.TextAlignTrailing, fyne.TextStyle{Monospace: true})
			return container.NewGridWithColumns(2, name, value)
		},
		func(id widget.ListItemID, o fyne.CanvasObject) {
			if id >= len(fanSDRs) {
				return
			}
			s := fanSDRs[id]
			grid := o.(*fyne.Container)

			valStr := "N/A"
			if s.Full != nil {
				valStr = fmt.Sprintf("%.0f RPM", s.Full.SensorValue)
			}

			grid.Objects[0].(*widget.Label).SetText(s.SensorName())
			grid.Objects[1].(*widget.Label).SetText(valStr)
		},
	)

	// 4028GR-TR
	fan4028Groups := container.NewVBox()
	zones := []string{"Zone 0:", "Zone 1:", "Zone 2:", "Zone 3:"}
	for i, name := range zones {
		zoneID := byte(i)
		slider := widget.NewSlider(0, 100)
		slider.SetValue(32) // default 0x20

		input := widget.NewEntry()
		input.SetText("32") // default 0x20
		valLabel := widget.NewLabel("Auto")
		slider.OnChanged = func(v float64) {
			input.SetText(fmt.Sprintf("%.0f", v))
		}
		input.OnChanged = func(s string) {
			var val int
			fmt.Sscanf(s, "%d", &val)
			if val >= 0 && val <= 100 {
				slider.SetValue(float64(val))
			}
		}

		// Set Speed
		setBtn := widget.NewButton("Set", func() {
			session.GetInstance().SetSuperMicroFanSpeed(zoneID, byte(slider.Value))
		})
		// auto Mode
		autoBtn := widget.NewButton("Auto", func() {
			session.GetInstance().SetSuperMicroFanSpeed(zoneID, 0xFF)
			valLabel.SetText("Auto (Released)")
		})

		hexPreview := widget.NewLabelWithStyle("", fyne.TextAlignTrailing, fyne.TextStyle{Monospace: true})
		go func(in *widget.Entry, lbl *widget.Label) {
			ticker := time.NewTicker(200 * time.Millisecond)
			for range ticker.C {
				fyne.Do(func() {
					var val int
					fmt.Sscanf(in.Text, "%d", &val)
					lbl.SetText(fmt.Sprintf("0x%02X", byte(val)))
				})
			}
		}(input, hexPreview)

		row := container.NewVBox(
			container.NewHBox(widget.NewLabel(name), layout.NewSpacer(), input, widget.NewLabel("%"), hexPreview),
			container.NewBorder(nil, nil, nil, container.NewHBox(setBtn, autoBtn), slider),
			widget.NewSeparator(),
		)
		fan4028Groups.Add(row)
	}
	modeRadio4028 := widget.NewRadioGroup([]string{"Full", "Optimal", "HeavyIO"}, func(s string) {
		switch s {
		case "Full":
			setContainerEnabled(fan4028Groups, true)
			session.GetInstance().SetSuperMicroFanFull()
		case "Optimal":
			setContainerEnabled(fan4028Groups, false)
			session.GetInstance().SetSuperMicroFanOptimal()
		case "HeavyIO":
			setContainerEnabled(fan4028Groups, false)
			session.GetInstance().SetSuperMicroFanHeavyIO()
		}
	})
	modeRadio4028.Horizontal = true

	fan4028Groups.Hide()
	modeRadio4028.Hide()
	// Select Server Type
	serverTypeSelect := widget.NewSelect([]string{"Generic", "SuperMicro 4028"}, func(value string) {
		fan4028Groups.Hide()
		modeRadio4028.Hide()
		if value == "SuperMicro 4028" {
			fan4028Groups.Show()
			setContainerEnabled(fan4028Groups, false)
			if modeRadio4028.Selected == "Full" {
				setContainerEnabled(fan4028Groups, true)
			}
			modeRadio4028.Show()
		}
	})
	serverTypeSelect.PlaceHolder = "Select Server Model"

	controlPanel := container.NewVBox(
		widget.NewCard("Global Settings", "Mode Switch",
			container.NewVBox(serverTypeSelect, modeRadio4028),
		),
		fan4028Groups,
	)

	go func() {
		ticker := time.NewTicker(1 * time.Second)
		for range ticker.C {
			fyne.Do(func() {
				allSDRs := session.GetInstance().GetSDRs()
				fanSDRs = getFanSDRs(allSDRs)
				fanRPMList.Refresh()
			})
		}
	}()

	split := container.NewHSplit(fanRPMList, container.NewPadded(container.NewVScroll(controlPanel)))
	split.Offset = 0.4
	return split
}

func getFanSDRs(allSDRs []*ipmi.SDR) []*ipmi.SDR {
	var temp []*ipmi.SDR
	for _, s := range allSDRs {
		if strings.Contains(strings.ToUpper(s.SensorName()), "FAN") {
			temp = append(temp, s)
		}
	}
	return temp
}
func setContainerEnabled(obj fyne.CanvasObject, enabled bool) {
	if d, ok := obj.(fyne.Disableable); ok {
		if enabled {
			d.Enable()
		} else {
			d.Disable()
		}
	}
	if c, ok := obj.(*fyne.Container); ok {
		for _, child := range c.Objects {
			setContainerEnabled(child, enabled)
		}
	}
}
