package ui

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func FansScreen(win fyne.Window) fyne.CanvasObject {

	loginButton := widget.NewButton("Form Dialog (Login Form)", func() {
		username := widget.NewEntry()
		username.Validator = validation.NewRegexp(`^[A-Za-z0-9_-]+$`, "username can only contain letters, numbers, '_', and '-'")
		password := widget.NewPasswordEntry()
		password.Validator = validation.NewRegexp(`^[A-Za-z0-9_-]+$`, "password can only contain letters, numbers, '_', and '-'")
		remember := false
		items := []*widget.FormItem{
			widget.NewFormItem("Username", username),
			widget.NewFormItem("Password", password),
			widget.NewFormItem("Remember me", widget.NewCheck("", func(checked bool) {
				remember = checked
			})),
		}
		dialog.ShowForm("Login...", "Log In", "Cancel", items, func(b bool) {
			if !b {
				return
			}
			var rememberText string
			if remember {
				rememberText = "and remember this login"
			}

			log.Println("Please Authenticate", username.Text, password.Text, rememberText)
		}, win)
	})
	return container.NewVScroll(container.NewVBox(loginButton))
}
