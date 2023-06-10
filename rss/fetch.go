package rss

import (
	"fmt"
	"html"

	"github.com/mmcdole/gofeed"
)

func FetchRss(s string) {
	fp := gofeed.NewParser()

	feed, err := fp.ParseURL(s)

	if err != nil {
		fmt.Println("Error fetching feed:", err)
		return
	}

	fmt.Println("Feed Title:", feed.Title)

	for _, item := range feed.Items {

		fmt.Println("Entry Title:", item.Title)

		fmt.Println("Published:", item.Published)

		content := html.UnescapeString(item.Description)
		cleanContent := StripHTMLTags(content)
		fmt.Println("Content:", cleanContent)

		for _, link := range item.Links {
			fmt.Println("Link:", link)
		}

		fmt.Println("--------------------------------------")
	}
}
