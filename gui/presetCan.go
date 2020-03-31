package main

import (
	"errors"
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
	createPresetBox := widget.NewButton("new", func() {
		o.canChan <- presetCreationCan(o)
	})
	back := widget.NewButton("back", func() {
		o.canChan <- homeCan(o)
	})
	v := widget.NewVBox(createPresetBox, back)
	return fyne.NewContainerWithLayout(layout.NewBorderLayout(nil, v, nil, nil), v, presetPreview(o, own, sub))
}

func presetPreview(o *orb, owner, sub []*types.Preset) *fyne.Container {
	header := fyne.NewContainerWithLayout(layout.NewGridLayout(12),
		widget.NewLabel("use/del"),
		widget.NewLabel("candidate"),
		widget.NewLabel("capitalization"),
		widget.NewLabel("color"),
		widget.NewLabel("discard"),
		widget.NewLabel("has link"),
		widget.NewLabel("make"),
		widget.NewLabel("odometer"),
		widget.NewLabel("price"),
		widget.NewLabel("required"),
		widget.NewLabel("subdomains"),
		widget.NewLabel("year"),
	)
	everyoneCon := fyne.NewContainerWithLayout(layout.NewGridLayout(12))
	everyoneBox := widget.NewGroup("default", everyoneCon)
	mineCon := fyne.NewContainerWithLayout(layout.NewGridLayout(12))
	mineBox := widget.NewGroup("mine", mineCon)
	sharedCon := fyne.NewContainerWithLayout(layout.NewGridLayout(12))
	sharedBox := widget.NewGroup("shared with me", sharedCon)
	vCon := widget.NewVBox(everyoneBox, mineBox, sharedBox)
	var con *fyne.Container
	for _, preset := range append(owner, sub...) {
		presetLabel := make([]*widget.Label, 11)
		suffix := ",\n"
		con = sharedCon
		if preset.Everyone {
			con = everyoneCon
		} else if o.username == preset.Owner {
			con = mineCon
		}
		h := widget.NewHBox()
		h.Append(widget.NewButton("use", func() {
			posts, err := getPosts(o, preset.Query)
			if err != nil {
				o.l.Fatalln(err.Error())
			}
			actual := make([]*types.Post, 0, len(posts))
			for _, p := range posts {
				if discardRequired(p, preset.Discard, preset.Required) {
					actual = append(actual, p)
				}
			}
			o.canChan <- postCan(o, actual, preset.Owner, 0, 50, presetCan)
		}))
		h.Append(widget.NewButton("del", func() {
			if err := deletePreset(o, preset.Id); err != nil {
				o.l.Fatalln(err.Error())
			}
			o.canChan <- presetCan(o)
		}))
		con.AddObject(h)
		for k := range preset.Query {
			switch k {
			case "candidate":
				presetLabel[0] = widget.NewLabel(strconv.FormatBool(preset.Candidate))
			case "cappercent":
				presetLabel[1] = widget.NewLabel(strconv.Itoa(preset.CapPercent))
			case "color":
				presetLabel[2] = widget.NewLabel(preset.Color)
			case "link":
				presetLabel[4] = widget.NewLabel(strconv.FormatBool(preset.Link))
			case "make":
				presetLabel[5] = widget.NewLabel(preset.Make)
			case "odometer":
				presetLabel[6] = widget.NewLabel(strconv.Itoa(preset.Odometer))
			case "price":
				presetLabel[7] = widget.NewLabel(strconv.Itoa(preset.Price))
			case "subdomain":
				subdomains := ""
				for _, s := range preset.SubDomains {
					subdomains += s + suffix
				}
				subdomains = strings.TrimSuffix(subdomains, suffix)
				presetLabel[9] = widget.NewLabel(subdomains)
			case "year":
				presetLabel[10] = widget.NewLabel(strconv.Itoa(preset.Year))
			default:
				o.l.Fatalln(errors.New("unexpected key in preset query: " + k))
			}
		}
		discards := ""
		for _, d := range preset.Discard {
			discards += d + suffix
		}
		discards = strings.TrimSuffix(discards, suffix)
		presetLabel[3] = widget.NewLabel(discards)
		require := ""
		for _, r := range preset.Required {
			require += r + suffix
		}
		require = strings.TrimSuffix(require, suffix)
		presetLabel[8] = widget.NewLabel(require)
		for _, v := range presetLabel {
			if v == nil {
				v = widget.NewLabel("")
			}
			con.AddObject(v)
		}
	}
	return fyne.NewContainerWithLayout(layout.NewBorderLayout(header, nil, nil, nil), header,
		widget.NewScrollContainer(vCon))
}
