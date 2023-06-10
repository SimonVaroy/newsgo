package main

import (
	"fmt"
	"html"

	"github.com/mmcdole/gofeed"
)

func main() {
	fp := gofeed.NewParser()

	feed, err := fp.ParseURL("https://lukesmith.xyz/index.xml")

	if err != nil {
		fmt.Println("Error fetching feed:", err)
		return
	}

	// Print feed title
	fmt.Println("Feed Title:", feed.Title)

	// Iterate over each feed entry
	for _, item := range feed.Items {
		// Print entry title
		fmt.Println("Entry Title:", item.Title)

		// Print publication date
		fmt.Println("Published:", item.Published)

		// Clean and print text content of the entry
		content := html.UnescapeString(item.Description)
		cleanContent := stripHTMLTags(content)
		fmt.Println("Content:", cleanContent)

		// Print any links present in the entry
		for _, link := range item.Links {
			fmt.Println("Link:", link)
		}

		fmt.Println("--------------------------------------")
	}
}

func stripHTMLTags(s string) string {
	var result string
	var withinTag bool

	for _, c := range s {
		if c == '<' {
			withinTag = true
		} else if c == '>' {
			withinTag = false
		} else if !withinTag {
			result += string(c)
		}
	}

	return result
}
