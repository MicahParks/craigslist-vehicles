package main

import (
	"strconv"
	"strings"

	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"

	"gitlab.com/MicahParks/cano-cars/types"
)

func presetCan(o *orb) *fyne.Container {
	own, sub, err := myPresets(o)
	if err != nil {
		o.l.Fatalln(err.Error())
	}
	return presetPreview(o, own, sub)
}

func presetPreview(o *orb, owner, sub []*types.Preset) *fyne.Container {
	header := fyne.NewContainerWithLayout(layout.NewGridLayout(12),
		widget.NewLabel("use"),
		widget.NewLabel("candidate"),
		widget.NewLabel("capitalization"),
		widget.NewLabel("color"),
		widget.NewLabel("discard"),
		widget.NewLabel("has link"),
		widget.NewLabel("make"),
		widget.NewLabel("odometer"),
		widget.NewLabel("price"),
		widget.NewLabel("required"),
		widget.NewLabel("shared with"),
		widget.NewLabel("subdomains"),
	)
	// TODO Build vCon from preset.Query.
	everyoneCon := fyne.NewContainerWithLayout(layout.NewGridLayout(12))
	everyoneBox := widget.NewGroup("default", everyoneCon)
	mineCon := fyne.NewContainerWithLayout(layout.NewGridLayout(12))
	mineBox := widget.NewGroup("mine", mineCon)
	sharedCon := fyne.NewContainerWithLayout(layout.NewGridLayout(12))
	sharedBox := widget.NewGroup("shared with me", sharedCon)
	vCon := widget.NewVBox(everyoneBox, mineBox, sharedBox)
	var con *fyne.Container
	for _, preset := range append(owner, sub...) {
		suffix := ",\n"
		discards := ""
		for _, d := range preset.Discard {
			discards += d + suffix
		}
		discards = strings.TrimSuffix(discards, suffix)
		require := ""
		for _, r := range preset.Required {
			require += r + suffix
		}
		require = strings.TrimSuffix(require, suffix)
		shares := ""
		for _, s := range preset.Subs {
			shares += s + suffix
		}
		shares = strings.TrimSuffix(shares, suffix)
		subdomains := ""
		for _, s := range preset.SubDomains {
			subdomains += s + suffix
		}
		subdomains = strings.TrimSuffix(subdomains, suffix)
		con = sharedCon
		if preset.Everyone {
			con = everyoneCon
		} else if o.username == preset.Owner {
			con = mineCon
		}
		con.AddObject(widget.NewButton("use", func() {
			posts, err := getPosts(o, preset.Query)
			if err != nil {
				o.l.Fatalln(err.Error())
			}
			o.canChan <- postCan(o, posts, preset.Owner, 0, 50, presetCan)
		}))
		con.AddObject(widget.NewLabel(strconv.FormatBool(preset.Candidate)))
		con.AddObject(widget.NewLabel(strconv.Itoa(preset.CapPercent)))
		con.AddObject(widget.NewLabel(preset.Color))
		con.AddObject(widget.NewLabel(discards))
		con.AddObject(widget.NewLabel(strconv.FormatBool(preset.Link)))
		con.AddObject(widget.NewLabel(preset.Make))
		con.AddObject(widget.NewLabel(strconv.Itoa(preset.Odometer)))
		con.AddObject(widget.NewLabel(strconv.Itoa(preset.Price)))
		con.AddObject(widget.NewLabel(require))
		con.AddObject(widget.NewLabel(shares))
		con.AddObject(widget.NewLabel(subdomains))
	}
	back := widget.NewButton("back", func() {
		o.canChan <- homeCan(o)
	})
	return fyne.NewContainerWithLayout(layout.NewBorderLayout(header, back, nil, nil), header, back,
		widget.NewScrollContainer(vCon))
}
