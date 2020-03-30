package types

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"unicode"

	"github.com/gocolly/colly/v2"
)

func (p *Post) GetAttr(e *colly.HTMLElement, l *log.Logger) {
	e.ForEach("p.attrgroup", func(_ int, el *colly.HTMLElement) {
		el.ForEach("span", func(_ int, elem *colly.HTMLElement) {
			text := strings.TrimSpace(elem.Text)
			split := strings.SplitN(text, ": ", 2)
			if len(split) == 2 {
				attribute := strings.TrimSpace(split[0])
				value := strings.TrimSpace(split[1])
				p.AttrGroup[attribute] = value
				if attribute == "odometer" {
					odo, err := strconv.Atoi(value)
					if err != nil {
						l.Fatalln(err.Error())
					}
					p.Odometer = odo
				}
			}
		})
	})
}

func (p *Post) GetCapPercent() {
	count := float64(0)
	total := float64(0)
	for _, v := range p.Title {
		if unicode.IsLetter(v) {
			total += 1
			if unicode.IsUpper(v) {
				count += 1
			}
		}
	}
	if count > 0 {
		p.CapPercent = int(count / total * 100)
	}
}

func (p *Post) GetColor(titleBody string) {
	colors := map[string]string{
		"yellow": "yellow",
		"red":    "red",
		"blue":   "blue",
		"navy":   "blue",
		"purple": "purple",
		"violet": "purple",
		"silver": "gray",
		"grey":   "gray",
		"gray":   "gray",
		"green":  "green",
		"white":  "white",
		"brown":  "brown",
		"black":  "black",
	}
	for k, v := range colors {
		re := regexp.MustCompile(fmt.Sprintf(`\b%s\b`, k))
		if found := re.FindString(titleBody); len(found) != 0 {
			p.Color = v
			break
		}
	}
}

func (p *Post) GetMake(titleBody string) {
	makers := map[string]string{
		"bmw":        "bmw",
		"mercedes":   "mercedes",
		"benz":       "mercedes",
		"dodge":      "dodge",
		"jeep":       "jeep",
		"ram":        "ram",
		"ford":       "ford",
		"lincoln":    "lincoln",
		"mazda":      "mazda",
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
		if found := re.FindString(titleBody); len(found) != 0 {
			p.Make = v
			break
		}
	}
}

func (p *Post) GetHasLink(titleBody string) {
	if strings.Contains(titleBody, ".com") ||
		strings.Contains(titleBody, "http") ||
		strings.Contains(titleBody, "www") {
		p.Link = true
	}
}

func (p *Post) GetYear(titleBody string) {
	re := regexp.MustCompile(`\b(((19)|(20))[0-9]{2})|(['"][0-9]{2})\b`)
	yearStr := re.FindString(titleBody)
	if strings.HasPrefix(yearStr, "'") || strings.HasPrefix(yearStr, `""`) {
		yearStr = "19" + yearStr[1:]
	}
	yearInt, err := strconv.Atoi(yearStr)
	if err != nil {
		return
	}
	p.Year = yearInt
}
