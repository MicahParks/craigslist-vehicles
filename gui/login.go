package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

func loginCanvas(o *orb) *widget.Box {

}

func registerCanvas(o *orb) fyne.CanvasObject {
	passwordBox := widget.NewPasswordEntry()
	passwordBox.SetPlaceHolder("password")
	usernameBox := widget.NewEntry()
	usernameBox.SetPlaceHolder("username")
	loginInstead := widget.NewButton("Already have a login?", func() {
		o.canvasC <- loginCanvas(o)
	})
	h := widget.NewHBox()
	return widget.NewVBox(
		widget.NewLabel("Please register with a unique username"),
		usernameBox,
		passwordBox,
		loginInstead,
	)
}
