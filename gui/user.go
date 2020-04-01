package main

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"

	"gitlab.com/MicahParks/craigslist-vehicles/types"
)

var (
	errAuth      = errors.New("incorrect password")
	errNotFound  = errors.New("user not found")
	errUserExist = errors.New("user already exists")
)

func allUsernames(o *orb) ([]string, error) {
	cursor, err := o.userCol.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	user := make([]*types.User, 0)
	if err = cursor.All(context.TODO(), &user); err != nil {
		return nil, err
	}
	names := make([]string, 0)
	for _, u := range user {
		names = append(names, u.Username)
	}
	return names, nil
}

func authenticate(o *orb, password, username string) error {
	user := &types.User{}
	auth := map[string]string{"_id": username}
	res := o.userCol.FindOne(context.TODO(), auth)
	if errors.Is(res.Err(), mongo.ErrNoDocuments) {
		return errNotFound
	}
	if err := res.Decode(user); err != nil {
		return err
	}
	if !checkPassword(user.Password, password) {
		return errAuth
	}
	o.username = user.Username
	o.user = user
	o.l.Printf("user %s authenticated", user.Username)
	return nil
}

func checkPassword(hash, password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		return false
	}
	return true
}

func getUser(o *orb) error {
	exist := map[string]string{"_id": o.username}
	cursor, err := o.userCol.Find(context.TODO(), exist)
	if err != nil {
		return err
	}
	user := make([]*types.User, 0)
	if err = cursor.All(context.TODO(), &user); err != nil {
		return err
	}
	o.user = user[0]
	return nil
}

func hashAndSalt(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func newUser(o *orb, password, username string) error {
	user := &types.User{}
	hash, err := hashAndSalt(password)
	if err != nil {
		o.l.Fatalln(err.Error())
	}
	user.Domains = []string{"richmond"}
	user.Username = username
	user.Password = hash
	exist := map[string]string{"_id": username}
	if res := o.userCol.FindOne(context.TODO(), exist); errors.Is(res.Err(), mongo.ErrNoDocuments) {
		if _, err = o.userCol.InsertOne(context.TODO(), user); err != nil {
			return err
		}
		o.username = user.Username
		o.user = user
		o.l.Printf("user %s created", user.Username)
		return nil
	}
	return errUserExist
}

func updateDomains(o *orb) error {
	update := bson.D{{"$set", bson.D{{"domains", o.user.Domains}}}}
	res, err := o.userCol.UpdateOne(context.TODO(), bson.M{"_id": o.username}, update)
	if err != nil {
		return err
	}
	if res.MatchedCount == 0 {
		return errors.New("user not found")
	}
	if res.UpsertedCount != 0 {
		return errors.New("created a new user instead of updating")
	}
	return nil
}
