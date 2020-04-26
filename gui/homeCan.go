package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

func homeCan(o *orb) *fyne.Container {
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
		o.canChan <- loginCan(o)
	})
	return fyne.NewContainerWithLayout(layout.NewBorderLayout(nil, logoutBox, nil, nil), logoutBox,
		widget.NewVBox(presetBox, domainBox, listBox))
}
