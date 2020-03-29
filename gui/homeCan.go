package main

import (
	"fyne.io/fyne/widget"
)

func homeCan(o *orb) *widget.Box {
	presetBox := widget.NewButton("presets", func() {
		o.canChan <- presetCan(o)
	})
	loginBox := widget.NewButton("logout", func() {
		o.username = ""
		o.user = nil
		// TODO Other logout stuff?
		o.canChan <- loginCan(o)
	})
	v := widget.NewVBox(presetBox, loginBox)
	return v
}
