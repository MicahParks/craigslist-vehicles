package main

import (
	"fyne.io/fyne/widget"

	"gitlab.com/MicahParks/cano-cars/types"
)

var (
	dontCare = ""
)

func presetCreationCan(o *orb) *widget.Form {
	p := &types.Preset{}
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

	candidateCheck := widget.NewCheck("", func(_ bool) {})
	candidateUse := widget.NewCheck("use", func(_ bool) {})
	cH := widget.NewHBox(candidateCheck, candidateUse)
	candidateF := widget.NewFormItem("candidate", cH)

	capPercentBox := widget.NewEntry()
	capPercentBox.SetPlaceHolder("0 through 100 (no percent)")
	capPercentF := widget.NewFormItem("capitalization percent", capPercentBox)

	colorBox := widget.NewSelect(colorOpts, func(_ string) {})
	colorBox.SetSelected(dontCare)
	colorF := widget.NewFormItem("color", colorBox)

	discardBox := widget.NewEntry()
	discardBox.SetPlaceHolder("discard1, discard2")
	discardF := widget.NewFormItem("discard", discardBox)

	linkCheck := widget.NewCheck("", func(_ bool) {})
	linkUse := widget.NewCheck("use", func(_ bool) {})
	lH := widget.NewHBox(linkCheck, linkUse)
	linkF := widget.NewFormItem("has link", lH)

	makeBox := widget.NewSelect(makeOpts, func(_ string) {})
	makeBox.SetSelected(dontCare)
	makeF := widget.NewFormItem("make", makeBox)

	odoBox := widget.NewEntry()
	odoBox.SetPlaceHolder("100000 (no commas)")
	odoF := widget.NewFormItem("odometer max", odoBox)

	priceBox := widget.NewEntry()
	priceBox.SetPlaceHolder("8000 (no commas or dollar signs)")
	priceF := widget.NewFormItem("price max", priceBox)

	requiredBox := widget.NewEntry()
	requiredBox.SetPlaceHolder("required1, required2")
	requiredF := widget.NewFormItem("required", requiredBox)

	subBox := widget.NewEntry()
	subBox.SetPlaceHolder("username, username2")
	subF := widget.NewFormItem("share", subBox)

	subdomainBox := widget.NewEntry()
	subdomainBox.SetPlaceHolder("richmond, washingtondc")
	subdomainF := widget.NewFormItem("subdomains", subdomainBox)

	yearBox := widget.NewEntry()
	yearBox.SetPlaceHolder("Ex. 1990")
	yearF := widget.NewFormItem("made after", yearBox)

	submit := widget.NewButton("create", func() {
		if err := p.MakeQuery(candidateCheck.Checked, candidateUse.Checked, capPercentBox.Text, colorBox.Selected, discardBox.Text,
			linkCheck.Checked, linkUse.Checked, makeBox.Selected, odoBox.Text, priceBox.Text, requiredBox.Text, subBox.Text,
			subdomainBox.Text, yearBox.Text); err != nil {
			o.l.Println(err.Error())
			return
		}
		p.Owner = o.username
		if err := insertPreset(o, p); err != nil {
			o.l.Fatalln(err.Error())
		}
		posts, err := getPosts(o, p.Query)
		if err != nil {
			o.l.Fatalln(err.Error())
		}
		o.canChan <- postCan(o, posts, 0, 50)
	})
	submitF := widget.NewFormItem("create", submit)
	return widget.NewForm(candidateF, capPercentF, colorF, discardF, linkF, makeF, odoF, priceF, requiredF, subF,
		subdomainF, yearF, submitF)
}
