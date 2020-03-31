package main

import (
	"errors"
	"strings"

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
		con.AddObject(widget.NewLabel(list.Name))
		if len(list.Posts) > 0 {
			query := bson.M{"_id": bson.M{"$in": list.Posts}}
			posts, err := getPosts(o, query)
			if err != nil {
				o.l.Fatalln(err.Error())
			}
			con.AddObject(widget.NewButton("view", func() {
				o.canChan <- postCan(o, posts, list.Owner, 0, 50, listCan)
			}))
		} else {
			b := widget.NewButton("nothing in list", func() {})
			b.Disable()
			con.AddObject(b)
		}
		del := widget.NewButton("delete", func() {
			if err = deleteList(o, list.Id); err != nil {
				o.l.Fatalln(err.Error())
			}
			o.canChan <- listCan(o)
		})
		if list.Owner != o.username {
			del.Disable()
		}
		con.AddObject(del)
	}
	e := widget.NewEntry()
	e.SetPlaceHolder("new list name")
	h := widget.NewHBox(e, widget.NewButton("add", func() {
		name := strings.TrimSpace(e.Text)
		if len(name) == 0 {
			o.l.Println("can't name list an empty string")
			return
		}
		if _, err := newList(o, name); err != nil {
			if errors.Is(err, errListExists) {
				o.l.Printf("list with name %s already exists", name)
				return
			}
			o.l.Fatalln(err.Error())
		}
		o.canChan <- listCan(o)
	}))
	v := widget.NewVBox(con, h)
	return fyne.NewContainerWithLayout(layout.NewMaxLayout(), v)
}
