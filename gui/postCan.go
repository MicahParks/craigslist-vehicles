package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"go.mongodb.org/mongo-driver/bson"
)

func postCan(o *orb) *fyne.Container {
	query := bson.M{"price": bson.M{"$lte": 5000}}
	posts, err := getPosts(o, query)
	if err != nil {
		o.l.Fatalln(err.Error())
	}
	boxes := rowVBoxes()
	for _, post := range posts {
		pR := postRow{post: post}
		if err := pR.make(); err != nil {
			o.l.Fatalln(err.Error())
		}
		pR.append(boxes)
	}
	con := fyne.NewContainerWithLayout(layout.NewGridLayout(8))
	for _, box := range boxes {
		con.AddObject(box)
	}
	return con
}
