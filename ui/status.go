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
	log.Println(sdrs==nil)
	for _, sdr := range sdrs {
		log.Printf(sdr.String())
	}
	content := container.NewVBox(
		widget.NewLabelWithStyle("\n\nSystem Status", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		widget.NewLabelWithStyle("\n\n", fyne.TextAlignCenter, fyne.TextStyle{Italic: true}))
	return content
}
