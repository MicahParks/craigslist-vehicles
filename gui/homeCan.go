package main

import (
	"fyne.io/fyne/widget"
)

func homeCan(o *orb) *widget.Box {
	presetBox := widget.NewButton("presets", func() {
		o.username = "test" // TODO Delete this.
		// TO
		o.canChan <- presetCan(o)
	})
	loginBox := widget.NewButton("logout", func() {
		o.username = ""
		o.user = nil
		// TODO Other logout stuff?
		o.canChan <- loginCan(o)
	})
	//tempBox := widget.NewButton("temp", func() {
	//	o.canChan <- postCan(o, 0, 50)
	//})
	return widget.NewVBox(presetBox, loginBox)
}
