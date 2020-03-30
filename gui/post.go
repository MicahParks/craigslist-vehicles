package main

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"gitlab.com/MicahParks/cano-cars/types"
)

func getPosts(o *orb, query bson.M, opts ...*options.FindOptions) (posts []*types.Post, err error) {
	posts = make([]*types.Post, 0)
	cursor, err := o.postsCol.Find(context.TODO(), query, opts...)
	if err != nil {
		return nil, err
	}
	if err = cursor.All(context.TODO(), &posts); err != nil {
		return nil, err
	}
	return
}

func updateCandidate(o *orb, post *types.Post) error {
	update := bson.D{{"$set", bson.D{{"candidate", true}}}}
	res, err := o.postsCol.UpdateOne(context.TODO(), bson.M{"_id": post.Url}, update)
	if err != nil {
		return err
	}
	if res.MatchedCount == 0 {
		return errors.New("post not found")
	}
	if res.UpsertedCount != 0 {
		return errors.New("created a new post instead of updating")
	}
	return nil
}
