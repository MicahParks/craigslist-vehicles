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

	"github.com/MicahParks/craigslist-vehicles/mongodb"
	"github.com/MicahParks/craigslist-vehicles/types"
)

var (
	Subdomains = []string{
		"frederick",
		"lynchburg",
		"norfolk",
		"richmond",
		"washingtondc",
	}
)

func main() {
	l := log.New(os.Stdout, "scraper: ", log.LstdFlags|log.LUTC|log.Lshortfile)
	collection, exists, err := mongodb.PostsExist()
	if err != nil {
		l.Println("Collection doesn't exist yet.")
	}
	if exists {
		if err = mongodb.DropCollection(collection); err != nil {
			l.Fatalln(err)
		}
	}
	for _, subdomain := range Subdomains {
		println(subdomain)
		domain := "%s.craigslist.org"
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
		if err = c.Limit(&colly.LimitRule{Delay: 0, DomainGlob: "*", Parallelism: 100}); err != nil {
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
			mux.Lock()
			post, ok := postUrl[url]
			if !ok {
				post = &types.Post{
					AttrGroup: make(map[string]string),
				}
				postUrl[url] = post
			}
			mux.Unlock()
			m := &types.Marsh{}
			if err = e.Unmarshal(m); err != nil {
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
				post.Subdomain = subdomain
			}()
		})
		if err = c.Visit(start); err != nil {
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
		if err = mongodb.InsertPosts(collection, posts); err != nil {
			l.Fatalln(err.Error())
		}
	}
}
