package main

import (
	"fyne.io/fyne/widget"
)

func homeCanvas(o *orb) *widget.Box {
	presetBox := widget.NewButton("presets", func() {
		o.canChan <- presetCanvas(o)
	})
	loginBox := widget.NewButton("logout", func() {
		o.username = ""
		o.user = nil
		// TODO Other logout stuff?
		o.canChan <- loginCanvas(o)
	})
	v := widget.NewVBox(presetBox, loginBox)
	return v
}
