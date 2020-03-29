package main

import (
	"fyne.io/fyne/widget"

	"gitlab.com/MicahParks/cano-cars/types"
)

var (
	dontCare = "don't care"
)

func presetCan(o *orb) *widget.Form {
	p := &types.Preset{}
	makeOpts := []string{
		dontCare,
		"acura",
		"bmw",
		"buick",
		"cadillac",
		"chevrolet",
		"dodge",
		"ford",
		"gm",
		"jeep",
		"kia",
		"lexus",
		"lincoln",
		"mazda",
		"mercedes",
		"honda",
		"hyundai",
		"nissan",
		"ram",
		"subaru",
		"toyota",
		"tesla",
		"volkswagen",
		"volvo",
	}
	colorOpts := []string{
		dontCare,
		"black",
		"blue",
		"brown",
		"gray",
		"green",
		"purple",
		"red",
		"white",
		"yellow",
	}
	colorBox := widget.NewSelect(colorOpts, func(color string) {
		p.Color = color
	})
	colorBox.SetSelected(dontCare)
	colorF := widget.NewFormItem("color", colorBox)
	makeBox := widget.NewSelect(makeOpts, func(make string) {
		p.Make = make
	})
	makeBox.SetSelected(dontCare)
	makeF := widget.NewFormItem("make", makeBox)
	modelBox := widget.NewEntry()
	modelBox.SetPlaceHolder("model or required string")
	modelF := widget.NewFormItem("unique", modelBox)
	shareBox := widget.NewEntry()
	shareBox.SetPlaceHolder("username, username2")
	shareF := widget.NewFormItem("share", shareBox)
	submit := widget.NewButton("create", func() {

	})
	submitF := widget.NewFormItem("create", submit)
	return widget.NewForm(colorF, makeF, modelF, shareF, submitF)
}
