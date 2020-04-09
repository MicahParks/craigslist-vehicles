package main

import (
	"context"
	"log"

	"gitlab.com/MicahParks/craigslist-vehicles/mongodb"
	"gitlab.com/MicahParks/craigslist-vehicles/types"
)

func main() {
	col, err := mongodb.Init("Preset")
	if err != nil {
		log.Fatalln(err.Error())
	}
	candidates := &types.Preset{}
	if err = candidates.MakeQuery("candidates", true, true, "", "", "",
		false, false, "", "", "", "", make([]string, 0),
		"richmond, lynchburg, frederick, washingtondc, norfolk", ""); err != nil {
		log.Fatalln(err.Error())
	}
	candidates.Everyone = true
	trucks := &types.Preset{}
	if err = trucks.MakeQuery("trucks", false, false, "", "", "", true,
		true, "ford", "80000", "20000", "", make([]string, 0),
		"richmond, lynchburg, frederick, washingtondc, norfolk", "2005"); err != nil {
		log.Fatalln(err.Error())
	}
	trucks.Everyone = true
	micahs := &types.Preset{}
	if err = micahs.MakeQuery("micahs", false, false, "50", "black",
		"", false, true, "", "100000", "10000", "", make([]string, 0),
		"richmond, lynchburg, frederick, washingtondc, norfolk", "2000"); err != nil {
		log.Fatalln(err.Error())
	}
	micahs.Everyone = true
	if _, err = col.InsertOne(context.TODO(), candidates); err != nil {
		log.Fatalln(err.Error())
	}
	if _, err = col.InsertOne(context.TODO(), trucks); err != nil {
		log.Fatalln(err.Error())
	}
	if _, err = col.InsertOne(context.TODO(), micahs); err != nil {
		log.Fatalln(err.Error())
	}
}
