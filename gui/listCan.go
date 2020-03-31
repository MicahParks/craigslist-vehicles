package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"go.mongodb.org/mongo-driver/bson"
)

func listCan(o *orb) *fyne.Container {
	lists, shared, err := myLists(o)
	if err != nil {
		o.l.Fatalln(err.Error())
	}
	con := fyne.NewContainerWithLayout(layout.NewGridLayout(3))
	for _, list := range append(lists, shared...) {
		b := bson.D{}
		for _, post := range list.Posts {
			b = bson.D{bson.E{Key: post, Value: post}}
		}
		query := bson.M{"$in": b}
		posts, err := getPosts(o, query)
		if err != nil {
			o.l.Fatalln(err.Error())
		}
		con.AddObject(widget.NewLabel(list.Name))
		con.AddObject(widget.NewButton("view", func() {
			o.canChan <- postCan(o, posts, list.Owner, 0, 50)
		}))
		del := widget.NewButton("delete", func() {
			if err = deleteList(o, list.Id); err != nil {
				o.l.Fatalln(err.Error())
			}
		})
		if list.Owner != o.username {
			del.Disable()
		}
		con.AddObject(del)
	}
	return con
}
