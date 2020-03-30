package types

import (
	"strconv"
	"strings"
)

func (p *Preset) Query(candidate bool, capPercent, color, discard string, link bool, makeCar,
	odometer, price, required, subs, subdomains string, year string) (post *Post, err error) {
	post = &Post{}
	var hold int
	p.Candidate = candidate
	post.Candidate = candidate
	if len(capPercent) != 0 {
		hold, err = strconv.Atoi(capPercent)
		if err != nil {
			return nil, err
		}
		p.CapPercent = hold
		post.CapPercent = hold
	}
	if len(color) != 0 {
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
	p.Link = link
	post.Link = link
	if len(makeCar) != 0 {
		p.Make = makeCar
		post.Make = makeCar
	}
	if len(odometer) != 0 {
		hold, err = strconv.Atoi(odometer)
		if err != nil {
			return nil, err
		}
		p.Odometer = hold
		post.Odometer = hold
	}
	if len(price) != 0 {
		hold, err = strconv.Atoi(price)
		if err != nil {
			return nil, err
		}
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
	if len(year) != 0 {
		hold, err = strconv.Atoi(year)
		if err != nil {
			return nil, err
		}
		p.Year = hold
		post.Year = hold
	}
	return post, nil
}
