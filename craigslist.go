package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"unicode"

	"github.com/gocolly/colly/v2"
)

// TODO Too many models. Make it a client side thing.

func (p *Post) attr(e *colly.HTMLElement, l *log.Logger) {
	e.ForEach("p.attrgroup", func(_ int, el *colly.HTMLElement) {
		el.ForEach("span", func(_ int, elem *colly.HTMLElement) {
			p.AttrGroup[el.Text] = elem.Text
			if el.Text == "odometer:" && len(elem.Text) > 0 {
				odo, err := strconv.Atoi(elem.Text)
				if err != nil {
					l.Fatalln(err.Error())
				}
				p.Odometer = odo
			}
		})
	})
}

func (p *Post) capPercent() {
	count := 0
	for _, v := range p.Title {
		if unicode.IsUpper(v) {
			count += 1
		}
	}
	p.CapPercent = len(p.Title) / count * 100
}

func (p *Post) color() {
	colors := map[string]string{
		"yellow": "yellow",
		"red":    "red",
		"blue":   "blue",
		"navy":   "blue",
		"purple": "purple",
		"violet": "purple",
		"silver": "silver",
		"grey":   "grey",
		"gray":   "gray",
		"green":  "green",
		"white":  "white",
		"brown":  "brown",
		"black":  "black",
	}
	for k, v := range colors {
		re := regexp.MustCompile(fmt.Sprintf(`\b%s\b`, k))
		if found := re.FindString(p.titleBody); len(found) != 0 {
			p.Color = v
			break
		}
	}
}

func (p *Post) getMake() {
	makers := map[string]string{
		"bmw":        "bmw",
		"mercedes":   "mercedes",
		"benz":       "mercedes",
		"dodge":      "dodge",
		"jeep":       "jeep",
		"ram":        "ram",
		"ford":       "ford",
		"lincoln":    "lincoln",
		"gm":         "gm",
		"gmc":        "gm",
		"buick":      "buick",
		"cadillac":   "cadillac",
		"chevy":      "chevrolet",
		"chevrolet":  "chevrolet",
		"acura":      "acura",
		"honda":      "honda",
		"hyundai":    "hyundai",
		"kai":        "kia",
		"nissan":     "nissan",
		"subaru":     "subaru",
		"lexus":      "lexus",
		"toyota":     "toyota",
		"tesla":      "tesla",
		"volkswagen": "volkswagen",
		"volvo":      "volvo",
	}
	for k, v := range makers {
		re := regexp.MustCompile(fmt.Sprintf(`\b%s\b`, k))
		if found := re.FindString(p.titleBody); len(found) != 0 {
			p.Make = v
			break
		}
	}
}

func (p *Post) hasLink() {
	if strings.Contains(p.titleBody, ".com") {
		p.HasLink = true
	}
}

func main() {
	l := log.New(os.Stdout, "cano cars scraper:", log.LstdFlags|log.LUTC|log.Lshortfile)
	// frederick, richmond, washingtondc
	domain := "%s.craigslist.org"
	subdomain := "richmond"
	domain = fmt.Sprintf(domain, subdomain)
	start := fmt.Sprintf("https://%s/search/cta?hasPic=1&bundleDuplicates=1", domain)
	collector := colly.NewCollector(
		colly.AllowedDomains(domain),
		colly.IgnoreRobotsTxt(),
	)
	posts := make(map[string]*Post)
	mux := &sync.Mutex{}
	collector.OnHTML("a.result-title.hdrlnk", func(e *colly.HTMLElement) {
		// Post tiles from query.
		// Grab the post's link.
		url := e.Attr("href")
		visited, err := collector.HasVisited(url)
		if err != nil {
			l.Fatalln(err.Error())
		}
		if visited {
			return
		}
		if err := e.Request.Visit(url); err != nil {
			l.Fatalf("error with URL: (%s) \"%s\"", url, err.Error())
		}
	})
	collector.OnHTML("a.button.next", func(e *colly.HTMLElement) {
		// Next button from query.
		// Follow it's link to request the next page.
		url := e.Attr("href")
		visited, err := collector.HasVisited(url)
		if err != nil {
			l.Fatalln(err.Error())
		}
		if visited {
			return
		}
		if err := e.Request.Visit(url); err != nil {
			l.Fatalf("error getting next page: \"%s\"", err.Error())
		}
	})
	collector.OnHTML("section.body", func(e *colly.HTMLElement) {
		// Post page.
		url := e.Request.URL.String()
		var post *Post
		mux.Lock()
		post, ok := posts[url]
		if !ok {
			post = &Post{}
			posts[url] = post
		}
		mux.Unlock()
		if err := e.Unmarshal(post); err != nil {
			l.Fatalln(err.Error())
		}
		post.titleBody = strings.ToLower(post.Title + "\n" + post.Text)
		post.attr(e, l)
		post.getMake()
		post.capPercent()
		post.year()
		post.Url = url
		post.Query = append(post.Query, domain)
	})
	if err := collector.Visit(start); err != nil {
		l.Fatalln(err.Error())
	}
	for _, v := range posts {
		// TODO Insert into mongo
	}
}

func (p *Post) year() {
	re := regexp.MustCompile(`\b(((19)|(20))[0-9]{2})|(['"][0-9]{2})\b`)
	yearStr := re.FindString(p.titleBody)
	yearInt, err := strconv.Atoi(yearStr)
	if err != nil {
		return
	}
	p.Year = yearInt
}
