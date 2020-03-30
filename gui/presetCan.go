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
	header := fyne.NewContainerWithLayout(layout.NewGridLayout(11),
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
	pCon := fyne.NewContainerWithLayout(layout.NewGridLayout(11))
	all := append(owner, sub...)
	for _, own := range all {
		suffix := ",\n"
		discards := ""
		for _, d := range own.Discard {
			discards += d + suffix
		}
		discards = strings.TrimSuffix(discards, suffix)
		require := ""
		for _, r := range own.Required {
			require += r + suffix
		}
		require = strings.TrimSuffix(require, suffix)
		shares := ""
		for _, s := range own.Subs {
			shares += s + suffix
		}
		shares = strings.TrimSuffix(shares, suffix)
		subdomains := ""
		for _, s := range own.SubDomains {
			subdomains += s + suffix
		}
		pCon.AddObject(widget.NewButton("use", func() {
			posts, err := getPosts(o, own.Query)
			if err != nil {
				o.l.Fatalln(err.Error())
			}
			o.canChan <- postCan(o, posts, 0, 50)
		}))
		subdomains = strings.TrimSuffix(subdomains, suffix)
		pCon.AddObject(widget.NewLabel(strconv.FormatBool(own.Candidate)))
		pCon.AddObject(widget.NewLabel(strconv.Itoa(own.CapPercent)))
		pCon.AddObject(widget.NewLabel(own.Color))
		pCon.AddObject(widget.NewLabel(discards))
		pCon.AddObject(widget.NewLabel(strconv.FormatBool(own.Link)))
		pCon.AddObject(widget.NewLabel(own.Make))
		pCon.AddObject(widget.NewLabel(strconv.Itoa(own.Odometer)))
		pCon.AddObject(widget.NewLabel(strconv.Itoa(own.Price)))
		pCon.AddObject(widget.NewLabel(require))
		pCon.AddObject(widget.NewLabel(shares))
		pCon.AddObject(widget.NewLabel(subdomains))
	}
	return fyne.NewContainerWithLayout(layout.NewBorderLayout(header, nil, nil, nil), header, pCon)
}
