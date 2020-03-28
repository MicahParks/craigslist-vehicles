package main

import (
	"log"
	"os"
)

func main() {
	l := log.New(os.Stdout, "cano cars scraper:", log.LstdFlags|log.LUTC|log.Lshortfile)
	start := "https://richmond.craigslist.org/search/cta?hasPic=1&bundleDuplicates=1"

}
