package main

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"gitlab.com/MicahParks/craigslist-vehicles/types"
)

var (
	errListExists = errors.New("list already exists")
)

func deleteList(o *orb, id string) error {
	if res := o.listCol.FindOneAndDelete(context.TODO(), map[string]string{"_id": id}); res.Err() != nil {
		return res.Err()
	}
	return nil
}

func getList(o *orb, name string) (*types.List, error) {
	exist := map[string]string{"owner": o.username, "name": name}
	cursor, err := o.listCol.Find(context.TODO(), exist)
	if err != nil {
		return nil, err
	}
	list := make([]*types.List, 0)
	if err = cursor.All(context.TODO(), &list); err != nil {
		return nil, err
	}
	return list[0], nil
}

func myLists(o *orb) (mine, shared []*types.List, err error) {
	mine = make([]*types.List, 0)
	shared = make([]*types.List, 0)
	ownQuery := bson.M{
		"owner": o.username,
	}
	cursor, err := o.listCol.Find(context.TODO(), ownQuery)
	if err != nil {
		return nil, nil, err
	}
	if err = cursor.All(context.TODO(), &mine); err != nil {
		return nil, nil, err
	}
	sharedQuery := bson.D{
		{"subs", o.username},
	}
	cursor, err = o.listCol.Find(context.TODO(), sharedQuery)
	if err != nil {
		return nil, nil, err
	}
	if err = cursor.All(context.TODO(), &shared); err != nil {
		return nil, nil, err
	}
	if mine == nil {
		mine = make([]*types.List, 0)
	}
	if shared == nil {
		shared = make([]*types.List, 0)
	}
	return mine, shared, nil
}

func newList(o *orb, name string) (*types.List, error) {
	list := &types.List{Id: o.username + name, Name: name, Owner: o.username, Subs: make([]string, 0)}
	if res := o.listCol.FindOne(context.TODO(), list); errors.Is(res.Err(), mongo.ErrNoDocuments) {
		if _, err := o.listCol.InsertOne(context.TODO(), list); err != nil {
			return nil, err
		}
		o.l.Printf("list %s created with %s as the owner", name, o.username)
		return getList(o, name)
	}
	return nil, errListExists
}

func updateList(o *orb, listId string, list *types.List) error {
	update := bson.M{"_id": listId}
	res := o.listCol.FindOneAndReplace(context.TODO(), update, list)
	if res.Err() != nil {
		return res.Err()
	}
	return nil
}
