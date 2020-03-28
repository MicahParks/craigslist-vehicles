package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly/v2"
)

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
	posts := make([]string, 0)
	collector.OnHTML(".result-title", func(e *colly.HTMLElement) {
		url := e.Attr("href")
		visited, err := collector.HasVisited(url)
		if err != nil {
			l.Fatalln(err.Error())
		}
		if visited {
			return
		}
		posts = append(posts, url)
	})
	collector.OnHTML(".next", func(e *colly.HTMLElement) {

	})
}
