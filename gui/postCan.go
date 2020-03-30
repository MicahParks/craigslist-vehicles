package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"

	"gitlab.com/MicahParks/cano-cars/types"
)

func postCan(o *orb, posts []*types.Post, start, end int) *fyne.Container {
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
	return fyne.NewContainerWithLayout(layout.NewBorderLayout(header, back, nil, nil), header, back, scroll)
}
