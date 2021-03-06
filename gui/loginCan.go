package main

import (
	"errors"

	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

func loginCan(o *orb) *widget.Box {
	headerLabel := widget.NewLabel("Please login.")
	passwordBox := widget.NewPasswordEntry()
	passwordBox.SetPlaceHolder("password")
	usernameBox := widget.NewEntry()
	usernameBox.SetPlaceHolder("username")
	loginInstead := widget.NewButton("Don't have a login?", func() {
		o.canChan <- registerCan(o)
	})
	submitBox := widget.NewButton("submit", func() {
		err := authenticate(o, passwordBox.Text, usernameBox.Text)
		if err != nil {
			if errors.Is(err, errAuth) || errors.Is(err, errNotFound) {
				o.l.Println(err.Error())
				headerLabel.SetText("Invalid login. Please try again.")
				return
			}
			o.l.Fatalln(err.Error())
		}
		o.canChan <- homeCan(o)
	})
	h := widget.NewHBox(loginInstead, submitBox)
	return widget.NewVBox(
		headerLabel,
		usernameBox,
		passwordBox,
		h,
	)
}

func registerCan(o *orb) fyne.CanvasObject {
	passwordBox := widget.NewPasswordEntry()
	passwordBox.SetPlaceHolder("password")
	usernameBox := widget.NewEntry()
	usernameBox.SetPlaceHolder("username")
	loginInstead := widget.NewButton("Already have a login?", func() {
		o.canChan <- loginCan(o)
	})
	submitBox := widget.NewButton("submit", func() {
		err := newUser(o, passwordBox.Text, usernameBox.Text)
		if err != nil {
			if errors.Is(err, errUserExist) {
				o.l.Println(err.Error())
				return
			}
			o.l.Fatalln(err.Error())
		}
		o.canChan <- homeCan(o)
	})
	h := widget.NewHBox(loginInstead, submitBox)
	return widget.NewVBox(
		widget.NewLabel("register"),
		usernameBox,
		passwordBox,
		h,
	)
}
