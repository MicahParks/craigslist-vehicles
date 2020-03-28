package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/gocolly/colly/v2"
)

// TODO Too many models. Make it a client side thing.

func main() {
	l := log.New(os.Stdout, "cano cars scraper: ", log.LstdFlags|log.LUTC|log.Lshortfile)
	collection, err := mongoInit()
	if err != nil {
		l.Fatalln(err.Error())
	}
	// frederick, richmond, washingtondc
	domain := "%s.craigslist.org"
	subdomain := "richmond"
	domain = fmt.Sprintf(domain, subdomain)
	start := fmt.Sprintf("https://%s/search/cta?hasPic=1&bundleDuplicates=1", domain)
	c := colly.NewCollector(
		colly.AllowedDomains(domain),
		colly.IgnoreRobotsTxt(),
		colly.DetectCharset(),
	)
	postUrl := make(map[string]*Post)
	mux := &sync.Mutex{}
	page := 0
	c.OnHTML("a.result-title.hdrlnk", func(e *colly.HTMLElement) {
		// Post tiles from query.
		// Grab the post's link.
		url := e.Attr("href")
		visited, err := c.HasVisited(url)
		if err != nil {
			l.Fatalln(err.Error())
		}
		if visited {
			return
		}
		go func() {
			if err := e.Request.Visit(url); err != nil {
				l.Fatalf("error with URL: (%s) \"%s\"", url, err.Error())
			}
		}()
	})
	c.OnHTML("a.button.next", func(e *colly.HTMLElement) {
		// Next button from query.
		// Follow it's link to request the next page.
		page = page + 1
		l.Printf("On page: %d have %d posts", page, len(postUrl))
		url := e.Attr("href")
		visited, err := c.HasVisited(url)
		if err != nil {
			l.Fatalln(err.Error())
		}
		if visited {
			return
		}
		if err := e.Request.Visit(url); err != nil {
			if err.Error() == "URL already visited" {
				return
			}
			l.Fatalf("error getting next page: \"%s\"", err.Error())
		}
	})
	c.OnHTML("section.body", func(e *colly.HTMLElement) {
		// Post page.
		url := e.Request.URL.String()
		var post *Post
		mux.Lock()
		post, ok := postUrl[url]
		if !ok {
			post = &Post{
				AttrGroup: make(map[string]string),
				Query:     make([]string, 0, 1),
			}
			postUrl[url] = post
		}
		mux.Unlock()
		m := &marsh{}
		if err := e.Unmarshal(m); err != nil {
			l.Fatalln(err.Error())
		}
		price, err := strconv.Atoi(strings.TrimPrefix(m.PriceStr, "$"))
		if err == nil {
			post.Price = price
		}
		post.Text = m.Text
		post.Title = m.Title
		post.titleBody = strings.ToLower(post.Title + "\n" + post.Text)
		post.attr(e, l)
		post.capPercent()
		post.color()
		post.getMake()
		post.hasLink()
		post.year()
		post.Url = url
		post.Query = append(post.Query, subdomain)
	})
	if err := c.Visit(start); err != nil {
		l.Fatalln(err.Error())
	}
	c.Wait()
	posts := make([]*Post, 0, len(postUrl))
	for _, v := range postUrl {
		posts = append(posts, v)
	}
	f, err := os.Create("scrape.gob")
	if err != nil {
		l.Fatalln(err.Error())
	}
	enc := gob.NewEncoder(f)
	if err = enc.Encode(posts); err != nil {
		l.Fatalln(err.Error())
	}
	if err = f.Close(); err != nil {
		l.Fatalln(err.Error())
	}
	if err := insertPosts(collection, posts); err != nil {
		l.Fatalln(err.Error())
	}
}
