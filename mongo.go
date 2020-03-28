package main

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func insertPosts(collection *mongo.Collection, posts []*Post) error {
	opts := options.InsertMany().SetOrdered(false)
	var many []interface{}
	for _, v := range posts {
		many = append(many, v)
	}
	_, err := collection.InsertMany(context.Background(), many, opts)
	if err != nil {
		return err
	}
	return nil
}

func mongoInit() (*mongo.Collection, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://0.0.0.0:27777/admin"))
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	return client.Database("cano").Collection("Posts"), nil
}
