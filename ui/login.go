package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func ShowLoginDialog(a fyne.App, w fyne.Window, stage *fyne.Container) {
	p := a.Preferences()
	host := widget.NewEntry()
	host.SetText(p.StringWithFallback("host", "127.0.0.1"))

	port := widget.NewEntry()
	port.SetText(p.StringWithFallback("port", "623"))

	user := widget.NewEntry()
	user.SetText(p.StringWithFallback("user", "ADMIN"))

	pass := widget.NewPasswordEntry()
	pass.SetText(p.StringWithFallback("pass", ""))

	items := []*widget.FormItem{
		{Text: "Host", Widget: host},
		{Text: "Port", Widget: port},
		{Text: "Username", Widget: user},
		{Text: "Password", Widget: pass},
	}

	d:=dialog.NewForm("IPMI Login","Connect","Exit",items,func (ok bool)  {
		if !ok{
			a.Quit()
			return
		}
		err:=
	})

}
