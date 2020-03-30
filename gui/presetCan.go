package main

import (
	"fyne.io/fyne"
)

func presetCan(o *orb) *fyne.Container {
	own, sub, err := myPresets(o)
	if err != nil {
		o.l.Fatalln(err.Error())
	}
	return presetPreview(o, own, sub)
}
