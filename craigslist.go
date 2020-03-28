package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"

	"github.com/gocolly/colly/v2"
)

func attr(e *colly.HTMLElement, l *log.Logger, post *Post) {
	e.ForEach("p.attrgroup", func(_ int, el *colly.HTMLElement) {
		el.ForEach("span", func(_ int, elem *colly.HTMLElement) {
			post.AttrGroup[el.Text] = elem.Text
			if el.Text == "odometer:" && len(elem.Text) > 0 {
				odo, err := strconv.Atoi(elem.Text)
				if err != nil {
					l.Fatalln(err.Error())
				}
				post.Odometer = odo
			}
		})
	})
}

func main() {
	l := log.New(os.Stdout, "cano cars scraper:", log.LstdFlags|log.LUTC|log.Lshortfile)
	// frederick, richmond, washingtondc
	domain := "%s.craigslist.org"
	subdomain := "richmond"
	domain = fmt.Sprintf(domain, subdomain)
	url := fmt.Sprintf("https://%s/search/cta?hasPic=1&bundleDuplicates=1", domain)
	collector := colly.NewCollector(
		colly.AllowedDomains(domain),
		colly.IgnoreRobotsTxt(),
	)
	posts := make(map[string]*Post)
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
	})
	collector.OnHTML("section.body", func(e *colly.HTMLElement) {
		url := e.Request.URL.String()
		var post *Post
		post, ok := posts[url]
		if !ok {
			post = &Post{}
			posts[url] = post
		}
		if err := e.Unmarshal(post); err != nil {
			l.Fatalln(err.Error())
		}
		titleBody := post.Title + post.Text
		attr(e, l, post)
		year(titleBody, post)
		post.Url = url
		post.Query = append(post.Query, domain)
	})
}

func year(titleBody string, post *Post) {
	re := regexp.MustCompile(`\b(((19)|(20))[0-9]{2})|(['"][0-9]{2})\b`)
	yearStr := re.FindString(titleBody)
	yearInt, err := strconv.Atoi(yearStr)
	if err != nil {
		return
	}
	post.Year = yearInt
}
