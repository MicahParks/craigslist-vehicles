package types

import (
	"strconv"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
)

func (p *Preset) MakeQuery(candidate, candidateUse bool, capPercent, color, discard string, link, linkUse bool, makeCar,
	odometer, price, required, subs, subDomains, year string) error {
	capPercent = strings.TrimSpace(capPercent)
	color = strings.TrimSpace(color)
	discard = strings.TrimSpace(discard)
	makeCar = strings.TrimSpace(makeCar)
	odometer = strings.TrimSpace(odometer)
	price = strings.TrimSpace(price)
	required = strings.TrimSpace(required)
	subs = strings.TrimSpace(subs)
	subDomains = strings.TrimSpace(subDomains)
	year = strings.TrimSpace(year)
	var hold int
	var err error
	post := &Post{}
	query := bson.M{}
	if candidateUse {
		query["candidate"] = candidate
		p.Candidate = candidate
		post.Candidate = candidate
	}
	if len(capPercent) != 0 {
		hold, err = strconv.Atoi(capPercent)
		if err != nil {
			return err
		}
		// TODO Add this to the query.
		p.CapPercent = hold
		post.CapPercent = hold
	}
	if len(color) != 0 {
		query["color"] = color
		p.Color = color
		post.Color = color
	}
	if len(discard) != 0 {
		p.Discard = make([]string, 0)
		for _, dis := range strings.Split(discard, ",") {
			dis = strings.TrimSpace(dis)
			if len(dis) != 0 {
				p.Discard = append(p.Subs, strings.TrimSpace(dis))
			}
		}
	}
	if linkUse {
		query["link"] = link
		p.Link = link
		post.Link = link
	}
	if len(makeCar) != 0 {
		query["make"] = makeCar
		p.Make = makeCar
		post.Make = makeCar
	}
	if len(odometer) != 0 {
		hold, err = strconv.Atoi(odometer)
		if err != nil {
			return err
		}
		// TODO Add this to the query.
		p.Odometer = hold
		post.Odometer = hold
	}
	if len(price) != 0 {
		hold, err = strconv.Atoi(price)
		if err != nil {
			return err
		}
		// TODO Add this to the query.
		p.Price = hold
		post.Price = hold
	}
	if len(required) != 0 {
		p.Required = make([]string, 0)
		for _, req := range strings.Split(required, ",") {
			req = strings.TrimSpace(req)
			if len(req) != 0 {
				p.Required = append(p.Required, strings.TrimSpace(req))
			}
		}
	}
	if len(subs) != 0 {
		p.Subs = make([]string, 0)
		for _, sub := range strings.Split(subs, ",") {
			sub = strings.TrimSpace(sub)
			if len(sub) != 0 {
				p.Subs = append(p.Subs, sub)
			}
		}
	}
	if len(subDomains) != 0 {
		p.SubDomains = make([]string, 0)
		for _, subD := range strings.Split(subDomains, ",") {
			subD = strings.TrimSpace(subD)
			if len(subD) != 0 {
				p.SubDomains = append(p.SubDomains, subD)
			}
		}
		// TODO Add this to the query.
	}
	if len(year) != 0 {
		hold, err = strconv.Atoi(year)
		if err != nil {
			return err
		}
		// TODO Add this to the query.
		p.Year = hold
		post.Year = hold
	}
	p.Query = query
	return nil
}
