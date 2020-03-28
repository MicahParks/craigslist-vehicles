package main

type Color string

type List struct {
	Id    string
	Owner string
	Subs  []string
}

type marsh struct {
	PriceStr string `selector:".price"`
	Text     string `selector:"#postingbody"`
	Title    string `selector:"#titletextonly"`
}

type Post struct {
	Url         string `bson:"_id"`
	CapPercent  int
	Color       string
	HasLink     bool
	IsCandidate bool
	Make        string
	Price       int
	Odometer    int
	titleBody   string
	Year        int
	Query       []string
	AttrGroup   map[string]string
	marsh
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

type Tile struct {
	Link string `selector:"li.result-row:nth-child(1) > p:nth-child(2) > a:nth-child(3)"`
}

type Query struct {
	Default bool
	Url     string
}
type User struct {
	Username string
	Password string
}
