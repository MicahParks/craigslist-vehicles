package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"

	"gitlab.com/MicahParks/cano-cars/types"
)

func postCan(o *orb, posts []*types.Post, start, end int) *fyne.Container {
	header := fyne.NewContainerWithLayout(layout.NewGridLayout(8))
	header.AddObject(widget.NewLabel("link"))
	header.AddObject(widget.NewLabel("price"))
	header.AddObject(widget.NewLabel("make"))
	header.AddObject(widget.NewLabel("odometer"))
	header.AddObject(widget.NewLabel("year"))
	header.AddObject(widget.NewLabel("color"))
	header.AddObject(widget.NewLabel("has link"))
	header.AddObject(widget.NewLabel("candidate"))
	boxes := rowVBoxes()
	for i, post := range posts {
		if i >= start {
			pR := postRow{post: post}
			if err := pR.make(); err != nil {
				o.l.Fatalln(err.Error())
			}
			pR.append(boxes)
		}
		if i >= end {
			break
		}
	}
	con := fyne.NewContainerWithLayout(layout.NewGridLayout(8))
	scroll := widget.NewScrollContainer(con)
	for _, box := range boxes {
		con.AddObject(box)
	}
	return fyne.NewContainerWithLayout(layout.NewBorderLayout(header, nil, nil, nil), header, scroll) // TODO Is this correct?
}
