package types

type Color string

type List struct {
	Id    string
	Owner string
	Subs  []string
}

type Marsh struct {
	PriceStr string `selector:".price"`
	Text     string `selector:"#postingbody"`
	Title    string `selector:"#titletextonly"`
}

type Post struct {
	Url         string `bson:"_id"`
	CapPercent  int
	Color       string
	HasLink     bool
	Text        string
	Title       string
	IsCandidate bool
	Make        string
	Price       int
	Odometer    int
	titleBody   string
	Year        int
	Query       []string
	AttrGroup   map[string]string
}

type Preset struct {
	Id      string
	Color   string
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
