package main

import (
	"context"

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
