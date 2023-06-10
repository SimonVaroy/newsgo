package rss

import (
	"fmt"
	"html"
	"os"

	"io/ioutil"
	"path/filepath"
	"strings"

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

func ReadFeedFile() ([]string, error) {
	homeDir := os.Getenv("HOME")
	filePath := filepath.Join(homeDir, ".config", "newsgo", "feeds")

	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error reading feeds file: %s", err)
	}

	lines := strings.Split(string(content), "\n")
	return lines, nil
}
