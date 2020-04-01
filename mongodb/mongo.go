package mongodb

import (
	"context"
	"io/ioutil"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"gitlab.com/MicahParks/craigslist-vehicles/types"
)

func InsertPosts(collection *mongo.Collection, posts []*types.Post) error {
	opts := options.InsertMany().SetOrdered(false)
	var many []interface{}
	for _, v := range posts {
		many = append(many, v)
	}
	if _, err := collection.InsertMany(context.TODO(), many, opts); err != nil {
		return err
	}
	return nil
}

func Init(collection string) (*mongo.Collection, error) {
	uriBytes, err := ioutil.ReadFile("mongo.uri")
	if err != nil {
		return nil, err
	}
	client, err := mongo.NewClient(options.Client().ApplyURI(strings.TrimSpace(string(uriBytes))))
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	return client.Database("cars").Collection(collection), nil
}
