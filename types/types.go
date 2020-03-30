package types

type List struct {
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
	Subdomain  []string
	Text       string
	Title      string
	Url        string `bson:"_id"`
	Year       int
	titleBody  string
}

type Preset struct {
	Everyone bool
	Owner    string
	Subs     []string

	Candidate  bool
	CapPercent int
	Color      string
	Discard    []string
	Link       bool
	Make       string
	Odometer   int
	Price      int
	Required   []string
	Subdomain  []string
	Year       int
}

type Query struct {
	Default bool
	Url     string `bson:"_id"`
}
type User struct {
	Username string `bson:"_id"`
	Password string
}
