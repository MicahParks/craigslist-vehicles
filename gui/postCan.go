package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"go.mongodb.org/mongo-driver/bson"
)

func postCan(o *orb, end, start int) *widget.ScrollContainer {
	query := bson.M{"price": bson.M{"$gte": 5000}}
	posts, err := getPosts(o, query)
	if err != nil {
		o.l.Fatalln(err.Error())
	}
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
