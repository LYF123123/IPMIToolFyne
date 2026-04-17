package ui

import (
	"IPMITOOLFYNE/session"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func StatusScreen(_ fyne.Window) fyne.CanvasObject {

	sdrs := session.GetInstance().GetSDRs()
	log.Println("is sdrs ==nil?")
	log.Println(sdrs == nil)
	var sdrLabel []fyne.CanvasObject
	sdrLabel = append(sdrLabel, widget.NewLabelWithStyle("\n\nSystem Status", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}))
	for _, sdr := range sdrs {
		// For now, I show every element from SDRs.
		// Of course, This is stupid, I will change soon
		log.Printf(sdr.String())
		sdrLabel = append(sdrLabel, widget.NewLabelWithStyle("\n\n"+sdr.String(), fyne.TextAlignCenter, fyne.TextStyle{Italic: true}))
	}
	content := container.NewVBox(sdrLabel...)
	scroll := container.NewVScroll(content)
	return scroll
}
