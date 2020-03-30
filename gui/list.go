package main

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"

	"gitlab.com/MicahParks/cano-cars/types"
)

func myLists(o *orb) ([]*types.List, error) {
	own := make([]*types.List, 0)
	ownQuery := bson.M{
		"owner": o.username,
	}
	cursor, err := o.presetCol.Find(context.TODO(), ownQuery)
	if err != nil {
		return nil, err
	}
	if err = cursor.All(context.TODO(), &own); err != nil {
		return nil, err
	}
	return own, nil
}

func updateList(o *orb, listId string, list *types.List) error {
	update := bson.M{"_id": listId}
	res, err := o.listCol.UpdateOne(context.TODO(), update, list)
	if err != nil {
		return err
	}
	if res.MatchedCount == 0 {
		return errors.New("list not found")
	}
	if res.UpsertedCount == 0 {
		return errors.New("created a new list instead of updating")
	}
	return nil
}
