package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"

	"gitlab.com/MicahParks/cano-cars/types"
)

func listAdd(o *orb, post *types.Post, back fyne.CanvasObject) {
	lists, err := myLists(o)
	if err != nil {
		o.l.Fatalln(err.Error())
	}
	form := widget.NewForm()
	for _, l := range lists {
		form.Append(l.Name, widget.NewButton("add", func() {
			for _, p := range l.Posts {
				if p == post.Url {
					o.canChan <- back
					return
				}
			}
			l.Posts = append(l.Posts, post.Url)
			if err = updateList(o, l.Id, l); err != nil {
				o.l.Fatalln(err.Error())
			}
			o.canChan <- back
		}))
	}

}
