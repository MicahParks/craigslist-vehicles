package main

import (
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"

	"gitlab.com/MicahParks/cano-cars/types"
)

func authenticate(o *orb, password, username string) (*types.User, error) {
	user := &types.User{}
	auth := map[string]string{"_id": username, "password": password}
	res := o.userCol.FindOne(context.TODO(), auth)
	if res == nil {
		return nil, errors.New("user not found")
	}
	if err := res.Decode(user); err != nil {
		return nil, err
	}
	return user, nil
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
