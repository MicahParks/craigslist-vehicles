package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

func homeCan(o *orb) *fyne.Container {
	// TODO Make listCan add button not suck and separate by ownership like preset.
	if err := getUser(o); err != nil {
		o.l.Fatalln(err.Error())
	}
	presetBox := widget.NewButton("presets", func() {
		o.canChan <- presetCan(o)
	})
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
	return fyne.NewContainerWithLayout(layout.NewBorderLayout(nil, logoutBox, nil, nil), logoutBox,
		widget.NewVBox(presetBox, domainBox, listBox))
}
