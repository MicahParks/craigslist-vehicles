package main

import (
	"context"
	"strconv"
	"strings"

	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"gitlab.com/MicahParks/cano-cars/types"
)

func insertPreset(o *orb, preset *types.Preset, opts ...*options.InsertOneOptions) error {
	_, err := o.presetCol.InsertOne(context.TODO(), preset, opts...)
	if err != nil {
		return err
	}
	return nil
}

func presetPreview(o *orb, owner, sub []*types.Preset) *fyne.Container {
	header := fyne.NewContainerWithLayout(layout.NewGridLayout(11),
		widget.NewLabel("candidate"),
		widget.NewLabel("capitalization"),
		widget.NewLabel("color"),
		widget.NewLabel("discard"),
		widget.NewLabel("has link"),
		widget.NewLabel("make"),
		widget.NewLabel("odometer"),
		widget.NewLabel("price"),
		widget.NewLabel("required"),
		widget.NewLabel("shared with"),
		widget.NewLabel("subdomains"),
	)
	pCon := fyne.NewContainerWithLayout(layout.NewGridLayout(11))
	all := append(owner, sub...)
	for _, own := range all {
		suffix := ",\n"
		discards := ""
		for _, d := range own.Discard {
			discards += d + suffix
		}
		discards = strings.TrimSuffix(discards, suffix)
		require := ""
		for _, r := range own.Required {
			require += r + suffix
		}
		require = strings.TrimSuffix(require, suffix)
		shares := ""
		for _, s := range own.Subs {
			shares += s + suffix
		}
		shares = strings.TrimSuffix(shares, suffix)
		subdomains := ""
		for _, s := range own.SubDomains {
			subdomains += s + suffix
		}
		click := func() {
			posts, err := getPosts(o, own.Query)
			if err != nil {
				o.l.Fatalln(err.Error())
			}
			o.canChan <- postCan(o, posts, 0, 50)
		}
		subdomains = strings.TrimSuffix(subdomains, suffix)
		pCon.AddObject(widget.NewButton(strconv.FormatBool(own.Candidate), click))
		pCon.AddObject(widget.NewButton(strconv.Itoa(own.CapPercent), click))
		pCon.AddObject(widget.NewButton(own.Color, click))
		pCon.AddObject(widget.NewButton(discards, click))
		pCon.AddObject(widget.NewButton(strconv.FormatBool(own.Link), click))
		pCon.AddObject(widget.NewButton(own.Make, click))
		pCon.AddObject(widget.NewButton(strconv.Itoa(own.Odometer), click))
		pCon.AddObject(widget.NewButton(strconv.Itoa(own.Price), click))
		pCon.AddObject(widget.NewButton(require, click))
		pCon.AddObject(widget.NewButton(shares, click))
		pCon.AddObject(widget.NewButton(subdomains, click))
	}
	return fyne.NewContainerWithLayout(layout.NewBorderLayout(header, nil, nil, nil), header, pCon)
}

func myPresets(o *orb) (own, shared []*types.Preset, err error) {
	shared = make([]*types.Preset, 0)
	own = make([]*types.Preset, 0)
	ownQuery := bson.M{
		"owner": o.username,
	}
	cursor, err := o.presetCol.Find(context.TODO(), ownQuery)
	if err != nil {
		return nil, nil, err
	}
	if err = cursor.All(context.TODO(), &own); err != nil {
		return nil, nil, err
	}
	sharedQuery := bson.D{
		{"subs", o.username},
	}
	cursor, err = o.presetCol.Find(context.TODO(), sharedQuery)
	if err != nil {
		return nil, nil, err
	}
	if err = cursor.All(context.TODO(), &shared); err != nil {
		return nil, nil, err
	}
	return own, shared, nil
}
