package main

import (
	"fyne.io/fyne/widget"

	"gitlab.com/MicahParks/cano-cars/types"
)

func login() (*widget.Box, *types.User) {

}

func register() *widget.Box {
	passwordBox := widget.NewPasswordEntry()
	passwordBox.SetPlaceHolder("password")
	usernameBox := widget.NewEntry()
	usernameBox.SetPlaceHolder("username")
	loginInstead :=
	h := widget.NewHBox()
	return widget.NewVBox(
		widget.NewLabel("Please register with a unique username"),
		usernameBox,
		passwordBox,
	)
}
