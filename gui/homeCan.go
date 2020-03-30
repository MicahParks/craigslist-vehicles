package main

import (
	"fyne.io/fyne/widget"
)

func homeCan(o *orb) *widget.Form {
	o.username = "test" // TODO Delete this.
	presetBox := widget.NewButton("view", func() {
		o.canChan <- presetCan(o)
	})
	createPresetBox := widget.NewButton("new", func() {
		o.canChan <- presetCreationCan(o)
	})
	hPreset := widget.NewHBox(presetBox, createPresetBox)
	loginBox := widget.NewButton("logout", func() {
		o.username = ""
		o.user = nil
		// TODO Other logout stuff?
		o.canChan <- loginCan(o)
	})
	//tempBox := widget.NewButton("temp", func() {
	//	o.canChan <- postCan(o, 0, 50)
	//})
	return widget.NewForm(widget.NewFormItem("preset", hPreset), widget.NewFormItem("logout", loginBox))
}
