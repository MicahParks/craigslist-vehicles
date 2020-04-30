package types

import (
	"go.mongodb.org/mongo-driver/bson"
)

type Domains struct {
	Default bool
	Domain  string `bson:"_id"`
}

type List struct {
	Id    string `bson:"_id"`
	Posts []string
	Name  string
	Owner string
	Subs  []string
}

type Marsh struct {
	PriceStr string `selector:".price"`
	Text     string `selector:"#postingbody"`
	Title    string `selector:"#titletextonly"`
}

type Post struct {
	AttrGroup  map[string]string
	Candidate  bool
	CapPercent int
	Color      string
	Link       bool
	Make       string
	Odometer   int
	Price      int
	Subdomain  string
	Text       string
	Title      string
	Url        string `bson:"_id"`
	Year       int
	titleBody  string
}

type Preset struct {
	Id       string `bson:"_id"`
	Everyone bool
	Owner    string
	Subs     []string
	Query    bson.M

	Candidate  bool
	CapPercent int
	Color      string
	Discard    []string
	Link       bool
	Make       string
	Odometer   int
	Price      int
	Required   []string
	SubDomains []string
	Year       int
}

type User struct {
	Deleted  []string
	Domains  []string
	Username string `bson:"_id"`
	Password string
}
