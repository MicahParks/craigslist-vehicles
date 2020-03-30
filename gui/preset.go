package main

import (
	"context"

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
