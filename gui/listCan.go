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
	ownedLists := fyne.NewContainerWithLayout(layout.NewGridLayout(4))
	ownedBox := widget.NewGroup("mine", ownedLists)
	sharedLists := fyne.NewContainerWithLayout(layout.NewGridLayout(4))
	sharedBox := widget.NewGroup("shared with me", sharedLists)
	groups := widget.NewVBox(ownedBox, sharedBox)
	for _, j := range append(lists, shared...) {
		list := j
		var con *fyne.Container
		con = sharedLists
		if list.Owner == o.username {
			con = ownedLists
		}
		con.AddObject(widget.NewLabel(list.Name))
		subBox := widget.NewButton("share", func() {
			userPop(o, &list.Subs, func() {
				if err = updateList(o, list.Id, list); err != nil {
					o.l.Fatalln(err.Error())
				}
			}).Show() // Lol the last four lines are crazy.
		})
		if o.username != list.Owner {
			subBox.Disable()
		}
		con.AddObject(subBox)
		if len(list.Posts) > 0 {
			query := bson.M{"_id": bson.M{"$in": list.Posts}}
			posts, err := getPosts(o, query)
			if err != nil {
				o.l.Fatalln(err.Error())
			}
			if list.Subs == nil {
				list.Subs = make([]string, 0)
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
		if _, err = newList(o, name); err != nil {
			if errors.Is(err, errListExists) {
				o.l.Printf("list with name %s already exists", name)
				return
			}
			o.l.Fatalln(err.Error())
		}
		o.canChan <- listCan(o)
	}))
	v := widget.NewVBox(groups, h)
	back := widget.NewButton("back", func() {
		o.canChan <- homeCan(o)
	})
	return fyne.NewContainerWithLayout(layout.NewBorderLayout(nil, back, nil, nil), back, v)
}
