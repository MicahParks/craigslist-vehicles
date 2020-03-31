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
	header := fyne.NewContainerWithLayout(layout.NewGridLayout(11),
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
		widget.NewLabel("subdomains"),
	)
	// TODO Build vCon from preset.Query.
	everyoneCon := fyne.NewContainerWithLayout(layout.NewGridLayout(11))
	everyoneBox := widget.NewGroup("default", everyoneCon)
	mineCon := fyne.NewContainerWithLayout(layout.NewGridLayout(11))
	mineBox := widget.NewGroup("mine", mineCon)
	sharedCon := fyne.NewContainerWithLayout(layout.NewGridLayout(11))
	sharedBox := widget.NewGroup("shared with me", sharedCon)
	vCon := widget.NewVBox(everyoneBox, mineBox, sharedBox)
	var con *fyne.Container
	for _, preset := range append(owner, sub...) {
		presetLabel := make([]*widget.Label, 10)
		suffix := ",\n"
		discards := ""
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
		for k, v := range preset.Query {
			switch k {
			case "candidate":
				a := v.(bool)
				presetLabel[0] = widget.NewLabel(strconv.Itoa(a))
			case "cappercent":
				a := v.(int)
				presetLabel[1] = widget.NewLabel(strconv.Itoa(a))
			case "color":
				a := v.(string)
				presetLabel[2] = widget.NewLabel(a)
			case "discard":
				a := v.([]string)
				for _, d := range a {
					discards += d + suffix
				}
				discards = strings.TrimSuffix(discards, suffix)
				presetLabel[3] = widget.NewLabel(discards)
			case "link":
				a := v.(bool)
				widget.NewLabel(strconv.FormatBool(a))
			case "make":
				a := v.(string)
				presetLabel[4] = widget.NewLabel(a)
			case "odometer":
				a := v.(int)
				presetLabel[5] = widget.NewLabel(strconv.Itoa(a))
			case "price":
				a := v.(int)
				presetLabel[6] = widget.NewLabel(strconv.Itoa(a))
			case "required":
				a := v.([]string)
				require := ""
				for _, r := range a {
					require += r + suffix
				}
				require = strings.TrimSuffix(require, suffix)
				presetLabel[7] = widget.NewLabel(require)
			case "subdomains":
				a := v.([]string)
				subdomains := ""
				for _, s := range a {
					subdomains += s + suffix
				}
				subdomains = strings.TrimSuffix(subdomains, suffix)
				presetLabel[8] = widget.NewLabel(subdomains)
			case "year":
				a := v.(int)
				presetLabel[9] = widget.NewLabel(strconv.Itoa(a))
			default:
				o.l.Fatalln(errors.New("unexpected key in preset query"))
			}
		}
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
