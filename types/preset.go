package types

import (
	"strconv"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
)

func (p *Preset) MakeQuery(id string, candidate, candidateUse bool, capPercent, color, discard string, link, linkUse bool, makeCar,
	odometer, price, required string, shares []string, subDomains, year string) error {
	p.Id = id
	capPercent = strings.TrimSpace(capPercent)
	color = strings.TrimSpace(color)
	discard = strings.TrimSpace(strings.ToLower(discard))
	makeCar = strings.TrimSpace(makeCar)
	odometer = strings.TrimSpace(odometer)
	price = strings.TrimSpace(price)
	required = strings.TrimSpace(strings.ToLower(required))
	subDomains = strings.TrimSpace(subDomains)
	year = strings.TrimSpace(year)
	var hold int
	var err error
	query := bson.M{}
	if candidateUse {
		query["candidate"] = candidate
		p.Candidate = candidate
	}
	if len(capPercent) != 0 {
		hold, err = strconv.Atoi(capPercent)
		if err != nil {
			return err
		}
		query["cappercent"] = bson.M{"$lte": hold}
		p.CapPercent = hold
	}
	if len(color) != 0 {
		query["color"] = color
		p.Color = color
	}
	if len(discard) != 0 {
		both := make([]string, 0, len(discard))
		p.Discard = make([]string, 0, len(discard))
		for _, dis := range strings.Split(discard, ",") {
			dis = strings.TrimSpace(dis)
			if len(dis) != 0 {
				both = append(both, dis)
			}
		}
		p.Discard = both
	}
	if linkUse {
		query["link"] = link
		p.Link = link
	}
	if len(makeCar) != 0 {
		query["make"] = makeCar
		p.Make = makeCar
	}
	if len(odometer) != 0 {
		hold, err = strconv.Atoi(odometer)
		if err != nil {
			return err
		}
		query["odometer"] = bson.M{"$lte": hold}
		p.Odometer = hold
	}
	if len(price) != 0 {
		hold, err = strconv.Atoi(price)
		if err != nil {
			return err
		}
		query["price"] = bson.M{"$lte": hold}
		p.Price = hold
	}
	if len(required) != 0 {
		both := make([]string, 0, len(required))
		p.Required = make([]string, 0, len(required))
		for _, req := range strings.Split(required, ",") {
			req = strings.TrimSpace(req)
			if len(req) != 0 {
				both = append(both, req)
			}
		}
		p.Required = both
	}
	p.Subs = shares
	if len(subDomains) != 0 {
		both := make([]string, 0, len(subDomains))
		p.SubDomains = make([]string, 0, len(subDomains))
		for _, subD := range strings.Split(subDomains, ",") {
			subD = strings.TrimSpace(subD)
			if len(subD) != 0 {
				both = append(both, subD)
			}
		}
		query["subdomain"] = bson.M{"$in": both}
		p.SubDomains = both
	}
	if len(year) != 0 {
		hold, err = strconv.Atoi(year)
		if err != nil {
			return err
		}
		query["year"] = bson.M{"$gte": year}
		p.Year = hold
	}
	p.Query = query
	return nil
}
