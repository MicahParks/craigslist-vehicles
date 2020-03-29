package main

import (
	"strconv"

	"fyne.io/fyne/widget"

	"gitlab.com/MicahParks/cano-cars/types"
)

var (
	dontCare = ""
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
	yearBox := widget.NewEntry()
	yearBox.SetPlaceHolder("must be made after this year")
	yearF := widget.NewFormItem("made after", yearBox)
	shareBox := widget.NewEntry()
	shareBox.SetPlaceHolder("username, username2")
	shareF := widget.NewFormItem("share", shareBox)
	submit := widget.NewButton("create", func() {
		year, err := strconv.Atoi(yearBox.Text)
		if err != nil {
			o.l.Println("couldn't convert year to integer")
			return
		}

		query := p.Query()

	})
	submitF := widget.NewFormItem("create", submit)
	return widget.NewForm(colorF, makeF, modelF, yearF, shareF, submitF)
}
