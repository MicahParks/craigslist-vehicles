package main

import (
	"errors"

	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

func loginCanvas(o *orb) *widget.Box {
	passwordBox := widget.NewPasswordEntry()
	passwordBox.SetPlaceHolder("password")
	usernameBox := widget.NewEntry()
	usernameBox.SetPlaceHolder("username")
	loginInstead := widget.NewButton("Don't have a login?", func() {
		o.canvasC <- registerCanvas(o)
	})
	submitBox := widget.NewButton("submit", func() {
		err := authenticate(o, passwordBox.Text, usernameBox.Text)
		if err != nil {
			if errors.Is(err, errAuth) || errors.Is(err, errNotFound) {
				// TODO Report error to the user.
				o.l.Println(err.Error())
				return
			}
			o.l.Fatalln(err.Error())
		}
	})
	h := widget.NewHBox(loginInstead, submitBox)
	return widget.NewVBox(
		widget.NewLabel("login"),
		usernameBox,
		passwordBox,
		h,
	)
}

func registerCanvas(o *orb) fyne.CanvasObject {
	passwordBox := widget.NewPasswordEntry()
	passwordBox.SetPlaceHolder("password")
	usernameBox := widget.NewEntry()
	usernameBox.SetPlaceHolder("username")
	loginInstead := widget.NewButton("Already have a login?", func() {
		o.canvasC <- loginCanvas(o)
	})
	submitBox := widget.NewButton("submit", func() {
		err := newUser(o, passwordBox.Text, usernameBox.Text)
		if err != nil {
			if errors.Is(err, errUserExist) {
				// TODO Report error to the user.
				o.l.Println(err.Error())
				return
			}
			o.l.Fatalln(err.Error())
		}
	})
	h := widget.NewHBox(loginInstead, submitBox)
	return widget.NewVBox(
		widget.NewLabel("register"),
		usernameBox,
		passwordBox,
		h,
	)
}