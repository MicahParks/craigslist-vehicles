package mongodb

import (
	"context"
	"io/ioutil"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/MicahParks/craigslist-vehicles/types"
)

func DropCollection(col *mongo.Collection) error {
	if _, err := col.DeleteMany(context.TODO(), bson.M{}); err != nil {
		return err
	}
	return nil
}

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

func PostsExist() (*mongo.Collection, bool, error) {
	uriBytes, err := ioutil.ReadFile("mongo.uri")
	if err != nil {
		return nil, false, err
	}
	client, err := mongo.NewClient(options.Client().ApplyURI(strings.TrimSpace(string(uriBytes))))
	if err != nil {
		return nil, false, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		return nil, false, err
	}
	names, err := client.Database("cars").ListCollectionNames(context.TODO(), bson.M{})
	if err != nil {
		return nil, false, err
	}
	for _, name := range names {
		if name == "Posts" {
			col, err := Init("Posts")
			if err != nil {
				return nil, true, err
			}
			return col, true, nil
		}
	}
	col, err := Init("Posts")
	if err != nil {
		return nil, false, err
	}
	return col, true, nil
}
