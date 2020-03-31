package main

import (
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"

	"gitlab.com/MicahParks/cano-cars/types"
)

func postCan(o *orb, posts []*types.Post, owner string, start, end int, backFun func(*orb) *fyne.Container) *fyne.Container {
	header := fyne.NewContainerWithLayout(layout.NewGridLayout(9),
		widget.NewLabel("link"),
		widget.NewLabel("price"),
		widget.NewLabel("make"),
		widget.NewLabel("odometer"),
		widget.NewLabel("year"),
		widget.NewLabel("color"),
		widget.NewLabel("has link"),
		widget.NewLabel("candidate"),
		widget.NewLabel("list"),
	)
	boxes := rowVBoxes()
	for i, post := range posts {
		if i >= start {
			pR := postRow{post: post}
			if err := pR.make(o); err != nil {
				o.l.Fatalln(err.Error())
			}
			pR.append(o, boxes, posts, owner, start, end)
		}
		if i >= end {
			break
		}
	}
	back := widget.NewButton("back", func() {
		o.canChan <- backFun(o)
	})
	con := fyne.NewContainerWithLayout(layout.NewGridLayout(9))
	scroll := widget.NewScrollContainer(con)
	for _, box := range boxes {
		con.AddObject(box)
	}
	info := widget.NewLabel(fmt.Sprintf("Owner: %s    Viewing %d - %d of %d", owner, start, end, len(posts)))
	left := widget.NewButton("<", func() {
		start = start - 50
		end = end - 50
		if start < 0 {
			start = 0
			end = 50
		}
		info.SetText(fmt.Sprintf("Owner: %s    Viewing %d - %d of %d", owner, start, end, len(posts)-1))
		o.canChan <- postCan(o, posts, owner, start, end, backFun)
	})
	right := widget.NewButton(">", func() {
		start = start + 50
		end = end + 50
		if end-1 > len(posts) {
			end = len(posts) - 1
			start = end - 1
		}
		info.SetText(fmt.Sprintf("Owner: %s    Viewing %d - %d of %d", owner, start, end, len(posts)-1))
		o.canChan <- postCan(o, posts, owner, start, end, backFun)
	})
	topH := widget.NewVBox(info, header)
	return fyne.NewContainerWithLayout(layout.NewBorderLayout(topH, back, left, right), topH, back, left, right, scroll)
}
