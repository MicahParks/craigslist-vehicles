package types

import (
	"go.mongodb.org/mongo-driver/bson"
)

func (p *Preset) Query() bson.M {
	query := bson.M{}
	if len(p.Color) != 0 {
		query["color"] = []bson.M{
			{"color": ""},
			{"color": p.Color},
		}
	}
	if len(p.Make) != 0 {
		query["make"] = p.Make
	}
	//if len(p.Model) != 0 {
	//
	//}
	if p.Year > 0 {
		query["year"] = p.Year
	}
	return query
}
