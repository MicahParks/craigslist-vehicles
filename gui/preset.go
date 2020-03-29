package main

import (
	"fyne.io/fyne/widget"

	"gitlab.com/MicahParks/cano-cars/types"
)

func preset() {
	var p *types.Preset
	options := []string{
		"yellow",
		"red",
		"blue",
		"blue",
		"purple",
		"purple",
		"gray",
		"gray",
		"gray",
		"green",
		"white",
		"brown",
		"black",
	}
	widget.NewSelect()
}
