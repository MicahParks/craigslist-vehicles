package main

import (
	"fyne.io/fyne/widget"
)

func homeCan(o *orb) *widget.Form {
	o.username = "test" // TODO Delete this.
	if err := getUser(o); err != nil {
		o.l.Fatalln(err.Error())
	}
	presetBox := widget.NewButton("view", func() {
		o.canChan <- presetCan(o)
	})
	createPresetBox := widget.NewButton("new", func() {
		o.canChan <- presetCreationCan(o)
	})
	hPreset := widget.NewHBox(presetBox, createPresetBox)
	domainBox := widget.NewButton("domains", func() {
		o.canChan <- domainsCan(o)
	})
	listBox := widget.NewButton("lists", func() {
		o.canChan <- listCan(o)
	})
	logoutBox := widget.NewButton("logout", func() {
		o.username = ""
		o.user = nil
		// TODO Other logout stuff?
		o.canChan <- loginCan(o)
	})
	return widget.NewForm(widget.NewFormItem("preset", hPreset), widget.NewFormItem("domains", domainBox),
		widget.NewFormItem("lists", listBox), widget.NewFormItem("logout", logoutBox))
}
