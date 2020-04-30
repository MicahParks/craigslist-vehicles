package main

import (
	"fmt"
	"regexp"
	"strings"

	"gitlab.com/MicahParks/craigslist-vehicles/types"
)

func discardRequired(o *orb, p *types.Post, discard []string, required []string) bool {
	if o.user.Deleted == nil {
		o.user.Deleted = make([]string, 0)
	} else {
		for _, url := range o.user.Deleted {
			if p.Url == url {
				return false
			}
		}
	}
	titleBody := p.Title + "\n" + p.Text
	titleBody = strings.ToLower(titleBody)
	for _, v := range discard {
		re := regexp.MustCompile(fmt.Sprintf(`\b%s\b`, v))
		if found := re.FindString(titleBody); len(found) != 0 {
			return false
		}
	}
	for _, v := range required {
		have := false
		re := regexp.MustCompile(fmt.Sprintf(`\b%s\b`, v))
		if found := re.FindString(titleBody); len(found) != 0 {
			have = true
		}
		if !have {
			return false
		}
	}
	return true
}
