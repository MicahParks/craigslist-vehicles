package main

import (
	"net/url"

	"fyne.io/fyne/widget"

	"gitlab.com/MicahParks/cano-cars/types"
)

type postRow struct {
	urlBox       *widget.Hyperlink
	priceBox     *widget.Label
	makeBox      *widget.Label
	odoBox       *widget.Label
	yearBox      *widget.Label
	colorBox     *widget.Label
	candidateBox *widget.Check
	titleBox     *widget.Label
	post         *types.Post
}

func rowVBoxes() []*widget.Box {
	return []*widget.Box{widget.NewVBox(), widget.NewVBox(), widget.NewVBox(), widget.NewVBox(), widget.NewVBox(),
		widget.NewVBox(), widget.NewVBox(), widget.NewVBox()}
}

func (p *postRow) append(boxes []*widget.Box) {
	boxes[0].Append(p.urlBox)
	boxes[1].Append(p.priceBox)
	boxes[2].Append(p.makeBox)
	boxes[3].Append(p.odoBox)
	boxes[4].Append(p.yearBox)
	boxes[5].Append(p.colorBox)
	boxes[6].Append(p.candidateBox)
	boxes[7].Append(p.titleBox)
}

func (p *postRow) attrBox() *widget.Form {
	form := widget.NewForm()
	for k, v := range p.post.AttrGroup {
		attr := widget.NewLabel(v)
		form.Append(k, attr)
	}
	return form
}

func (p *postRow) make() error {
	u, err := url.Parse(p.post.Url)
	if err != nil {
		return err
	}
	p.urlBox = widget.NewHyperlink("link", u)
	p.priceBox = widget.NewLabel("$" + string(p.post.Price))
	p.makeBox = widget.NewLabel(p.post.Make)
	p.odoBox = widget.NewLabel(string(p.post.Odometer))
	p.yearBox = widget.NewLabel(string(p.post.Year))
	p.colorBox = widget.NewLabel(p.post.Color)
	p.candidateBox = widget.NewCheck("", func(_ bool) {})
	p.candidateBox.Checked = p.post.IsCandidate
	p.candidateBox.Disable()
	p.titleBox = widget.NewLabel(p.post.Title)
	return nil
}
