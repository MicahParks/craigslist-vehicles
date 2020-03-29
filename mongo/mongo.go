package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"gitlab.com/MicahParks/cano-cars/types"
)

func InsertPosts(collection *mongo.Collection, posts []*types.Post) error {
	opts := options.InsertMany().SetOrdered(false)
	var many []interface{}
	for _, v := range posts {
		many = append(many, v)
	}
	_, err := collection.InsertMany(context.TODO(), many, opts)
	if err != nil {
		return err
	}
	return nil
}

func Init() (*mongo.Collection, error) {
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
