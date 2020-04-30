package main

import (
	"fmt"
	"net/url"

	"fyne.io/fyne/widget"

	"gitlab.com/MicahParks/craigslist-vehicles/types"
)

type postRow struct {
	urlBox        *widget.Hyperlink
	priceBox      *widget.Label
	makeBox       *widget.Label
	odoBox        *widget.Label
	yearBox       *widget.Label
	colorBox      *widget.Label
	linkBox       *widget.Check
	candidateBox  *widget.Check
	candidateButt *widget.Button
	titleBox      *widget.Label
	post          *types.Post
}

func rowVBoxes() []*widget.Box {
	urlBox := widget.NewVBox()
	priceBox := widget.NewVBox()
	makeBox := widget.NewVBox()
	odoBox := widget.NewVBox()
	yearBox := widget.NewVBox()
	colorBox := widget.NewVBox()
	linkBox := widget.NewVBox()
	candidateBox := widget.NewVBox()
	listBox := widget.NewVBox()
	return []*widget.Box{urlBox, priceBox, makeBox, odoBox, yearBox, colorBox, linkBox, candidateBox, listBox}
}

func (p *postRow) append(o *orb, boxes []*widget.Box, posts []*types.Post, owner string, start, end int) {
	delButt := widget.NewButton("del", func() {
		if o.user.Deleted == nil {
			o.user.Deleted = make([]string, 0)
		}
		o.user.Deleted = append(o.user.Deleted, p.post.Url)
		if err := updateUser(o); err != nil {
			o.l.Fatalln(err.Error())
		}
		newPosts := make([]*types.Post, 0, len(posts)-1)
		for _, post := range posts {
			if post != p.post {
				newPosts = append(newPosts, post)
			}
		}
		o.canChan <- postCan(o, newPosts, o.username, start, end, presetCan)
	})
	boxes[0].Append(widget.NewHBox(p.urlBox, delButt))
	boxes[1].Append(p.priceBox)
	boxes[2].Append(p.makeBox)
	boxes[3].Append(p.odoBox)
	boxes[4].Append(p.yearBox)
	boxes[5].Append(p.colorBox)
	boxes[6].Append(p.linkBox)
	boxes[7].Append(widget.NewHBox(p.candidateBox, p.candidateButt))
	boxes[8].Append(widget.NewButton("add", func() {
		o.canChan <- listAdd(o, p.post, posts, owner, start, end)
	}))
}

func (p *postRow) attrBox() *widget.Form {
	form := widget.NewForm()
	for k, v := range p.post.AttrGroup {
		attr := widget.NewLabel(v)
		form.Append(k, attr)
	}
	return form
}

func (p *postRow) make(o *orb, candidateList *types.List) error {
	u, err := url.Parse(p.post.Url)
	if err != nil {
		return err
	}
	p.urlBox = widget.NewHyperlink("link", u)

	priceStr := "?"
	if p.post.Price > 0 {
		priceStr = fmt.Sprintf("$%d", p.post.Price)
	}
	p.priceBox = widget.NewLabel(priceStr)

	if len(p.post.Make) == 0 {
		p.post.Make = "?"
	}
	p.makeBox = widget.NewLabel(p.post.Make)

	odoStr := "?"
	if p.post.Odometer > 0 {
		odoStr = fmt.Sprintf("%d", p.post.Odometer)
	}
	p.odoBox = widget.NewLabel(odoStr)

	yearStr := "?"
	if p.post.Year > 0 {
		yearStr = fmt.Sprintf("%d", p.post.Year)
	}
	p.yearBox = widget.NewLabel(yearStr)

	if len(p.post.Color) == 0 {
		p.post.Color = "?"
	}
	p.colorBox = widget.NewLabel(p.post.Color)

	p.linkBox = widget.NewCheck("", func(_ bool) {})
	p.linkBox.Checked = p.post.Link
	p.linkBox.Disable()

	p.candidateBox = widget.NewCheck("", func(_ bool) {})
	p.candidateBox.Checked = p.post.Candidate
	p.candidateBox.Disable()

	p.candidateButt = widget.NewButton("mark", func() {
		if !p.post.Candidate {
			p.candidateBox.Checked = true
			p.candidateBox.Refresh()
			p.post.Candidate = true
			if err = updateCandidate(o, p.post); err != nil {
				o.l.Fatalln(err.Error())
			}
		}
		already := false
		for _, post := range candidateList.Posts {
			if post == p.post.Url {
				already = true
				break
			}
		}
		if !already {
			candidateList.Posts = append(candidateList.Posts, p.post.Url)
			if err = updateList(o, candidateList.Id, candidateList); err != nil {
				o.l.Fatalln(err.Error())
			}
		}
		p.candidateButt.Disable()
	})
	for _, post := range candidateList.Posts {
		if post == p.post.Url {
			p.candidateButt.Disable()
		}
	}

	p.titleBox = widget.NewLabel(p.post.Title)
	return nil
}
