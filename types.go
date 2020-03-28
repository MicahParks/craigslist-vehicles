package main

type User struct {
	Username string
	Password string
}

type Query struct {
	Url  string
	User string
}

type Post struct {
	Url         string
	CapPercent  int
	Color       string
	HasLink     bool
	IsCandidate bool
	Make        string
	Model       string
	Odometer    int
	Price       int
	Text        string
	Title       string
	Year        int
}

type QueryPost struct {
	Id    string
	Query string
	Post  string
}

type UPost struct {
	Id   string
	Post string
	User string
	Type string
}

type Preset struct {
	Id    string
	Color string // Make another type predefined?
	Make  string
	Model string
	Year  int
}

type PShare struct {
	Id    string
	Owner string
}

type PShareSubs struct {
	Id   string
	User string
}

type UserPreset struct {
	Preset string
	User   string
}

type List struct {
	Id    string
	Owner string
}

type ListPosts struct {
	Id   string
	Post string
}

type ListSubs struct {
	Id   string
	User string
}
