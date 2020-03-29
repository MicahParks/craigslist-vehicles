package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"

	"gitlab.com/MicahParks/cano-cars/types"
)

func postCan(o *orb, posts []*types.Post, start, end int) *widget.ScrollContainer {
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
	return scroll
}
