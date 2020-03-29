package main

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"

	"gitlab.com/MicahParks/cano-cars/types"
)

var (
	errAuth      = errors.New("incorrect password")
	errNotFound  = errors.New("user not found")
	errUserExist = errors.New("user already exists")
)

func authenticate(o *orb, password, username string) error {
	user := &types.User{}
	auth := map[string]string{"_id": username}
	res := o.userCol.FindOne(context.TODO(), auth)
	if res == nil {
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
	return nil
}

func checkPassword(hash, password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		return false
	}
	return true
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
	user.Password = password
	user.Username = username
	exist := map[string]string{"_id": username}
	if res := o.userCol.FindOne(context.TODO(), exist); errors.Is(res.Err(), mongo.ErrNoDocuments) {
		if _, err := o.userCol.InsertOne(context.TODO(), user); err != nil {
			return err
		}
		return nil
	}
	return errUserExist
}
