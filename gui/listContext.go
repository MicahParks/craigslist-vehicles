package main

import (
	"fyne.io/fyne/widget"

	"github.com/MicahParks/craigslist-vehicles/types"
)

func listAdd(o *orb, post *types.Post, posts []*types.Post, owner string, start, end int) *widget.Form {
	lists, _, err := myLists(o)
	if err != nil {
		o.l.Fatalln(err.Error())
	}
	if len(lists) == 0 {
		l, err := newList(o, "default")
		if err != nil {
			o.l.Fatalln(err.Error())
		}
		lists = []*types.List{l}
	}
	form := widget.NewForm()
	back := postCan(o, posts, owner, start, end, presetCan)
	for _, l := range lists {
		l := l
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
			if l.Id == o.username+"Candidates" {
				post.Candidate = true
				if err = updateCandidate(o, post); err != nil {
					o.l.Fatalln(err.Error())
				}
			}
			o.canChan <- back
		}))
	}
	return form
}
