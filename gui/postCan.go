package main

import (
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"

	"gitlab.com/MicahParks/cano-cars/types"
)

func postCan(o *orb, posts []*types.Post, preset *types.Preset, start, end int) *fyne.Container {
	header := fyne.NewContainerWithLayout(layout.NewGridLayout(8),
		widget.NewLabel("link"),
		widget.NewLabel("price"),
		widget.NewLabel("make"),
		widget.NewLabel("odometer"),
		widget.NewLabel("year"),
		widget.NewLabel("color"),
		widget.NewLabel("has link"),
		widget.NewLabel("candidate"),
	)
	boxes := rowVBoxes()
	for i, post := range posts {
		if i >= start {
			pR := postRow{post: post}
			if err := pR.make(o); err != nil {
				o.l.Fatalln(err.Error())
			}
			pR.append(boxes)
		}
		if i >= end {
			break
		}
	}
	back := widget.NewButton("back", func() {
		o.canChan <- presetCan(o)
	})
	con := fyne.NewContainerWithLayout(layout.NewGridLayout(8))
	scroll := widget.NewScrollContainer(con)
	for _, box := range boxes {
		con.AddObject(box)
	}
	info := widget.NewLabel(fmt.Sprintf("Owner: %s    Viewing %d - %d of %d", preset.Owner, start, end, len(posts)))
	left := widget.NewButton("<", func() {
		start = start - 50
		end = end - 50
		if start < 0 {
			start = 0
			end = 50
		}
		info.SetText(fmt.Sprintf("Owner: %s    Viewing %d - %d of %d", preset.Owner, start, end, len(posts)-1))
		o.canChan <- postCan(o, posts, preset, start, end)
	})
	right := widget.NewButton(">", func() {
		start = start + 50
		end = end + 50
		if end-1 > len(posts) {
			end = len(posts) - 1
			start = end - 1
		}
		info.SetText(fmt.Sprintf("Owner: %s    Viewing %d - %d of %d", preset.Owner, start, end, len(posts)-1))
		o.canChan <- postCan(o, posts, preset, start, end)
	})
	topH := widget.NewVBox(info, header)
	return fyne.NewContainerWithLayout(layout.NewBorderLayout(topH, back, left, right), topH, back, left, right, scroll)
}
