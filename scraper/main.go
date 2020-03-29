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

	"gitlab.com/MicahParks/cano-cars/mongodb"
	"gitlab.com/MicahParks/cano-cars/types"
)

// TODO Too many models. Make it a client side thing.

func main() {
	l := log.New(os.Stdout, "cano cars scraper: ", log.LstdFlags|log.LUTC|log.Lshortfile)
	collection, err := mongodb.Init("Posts")
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
		colly.Async(true),
	)
	postUrl := make(map[string]*types.Post)
	mux := &sync.Mutex{}
	page := 0
	wg := &sync.WaitGroup{}
	if err = c.Limit(&colly.LimitRule{Delay: 0, DomainGlob: "*", Parallelism: 55}); err != nil {
		l.Fatalln(err.Error())
	}
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
		wg.Add(1)
		go func() {
			wg.Done()
			if err := e.Request.Visit(url); err != nil {
				if err.Error() == "URL already visited" {
					return
				}
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
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := e.Request.Visit(url); err != nil {
				if err.Error() == "URL already visited" {
					return
				}
				l.Fatalf("error getting next page: \"%s\"", err.Error())
			}
		}()
	})
	c.OnHTML("section.body", func(e *colly.HTMLElement) {
		// Post page.
		url := e.Request.URL.String()
		var post *types.Post
		mux.Lock()
		post, ok := postUrl[url]
		if !ok {
			post = &types.Post{
				AttrGroup: make(map[string]string),
				Query:     make([]string, 0, 1),
			}
			postUrl[url] = post
		}
		mux.Unlock()
		m := &types.Marsh{}
		if err := e.Unmarshal(m); err != nil {
			l.Fatalln(err.Error())
		}
		wg.Add(1)
		go func() {
			defer wg.Done()
			price, err := strconv.Atoi(strings.TrimPrefix(m.PriceStr, "$"))
			if err == nil {
				post.Price = price
			}
			post.Text = m.Text
			post.Title = m.Title
			titleBody := strings.ToLower(post.Title + "\n" + post.Text)
			post.GetAttr(e, l)
			post.GetCapPercent()
			post.GetColor(titleBody)
			post.GetMake(titleBody)
			post.GetHasLink(titleBody)
			post.GetYear(titleBody)
			post.Url = url
			post.Query = append(post.Query, subdomain)
		}()
	})
	if err := c.Visit(start); err != nil {
		l.Fatalln(err.Error())
	}
	wg.Wait()
	c.Wait()
	posts := make([]*types.Post, 0, len(postUrl))
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
	if err := mongodb.InsertPosts(collection, posts); err != nil {
		l.Fatalln(err.Error())
	}
}
