package main

type Color string

type List struct {
	Id    string
	Owner string
	Subs  []string
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
	Query       Query
}

type Preset struct {
	Id      string
	Color   Color
	Default bool
	Make    string
	Model   string
	Owner   string
	Subs    []string
	Year    int
}

type Query struct {
	Default bool
	Url     string
}
type User struct {
	Username string
	Password string
}
