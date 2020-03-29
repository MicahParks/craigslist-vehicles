package main

import (
	"fyne.io/fyne/widget"

	"gitlab.com/MicahParks/cano-cars/types"
)

var (
	dontCare = "don't care"
)

func preset() *widget.Form {
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
	colorF := widget.NewFormItem("Color", colorBox)
	makeBox := widget.NewSelect(makeOpts, func(make string) {
		p.Make = make
	})
	makeBox.SetSelected(dontCare)
	makeF := widget.NewFormItem("Make", makeBox)
	modelBox := widget.NewEntry()
	modelBox.SetPlaceHolder("model or required string")
	modelF := widget.NewFormItem("Unique", modelBox)
	shareBox := widget.NewEntry()
	shareBox.SetPlaceHolder("username, username2")
	shareF := widget.NewFormItem("Share", shareBox)
	return widget.NewForm(colorF, makeF, modelF, shareF)
}
