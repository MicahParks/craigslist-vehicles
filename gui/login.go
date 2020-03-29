package main

import (
	"fyne.io/fyne/widget"
)

func register() *widget.Box {
	passwordBox := widget.NewPasswordEntry()
	passwordBox.SetPlaceHolder("password")
	usernameBox := widget.NewEntry()
	usernameBox.SetPlaceHolder("username")
	return widget.NewVBox(
		widget.NewLabel("Please register with a unique username"),
		usernameBox,
		passwordBox,
	)
}
