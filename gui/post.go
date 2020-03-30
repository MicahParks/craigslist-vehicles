package main

import (
	"context"

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
