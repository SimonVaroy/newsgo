package main

import (
	"fmt"

	"github.com/SimonVaroy/newsgo/filesystem"
	"github.com/SimonVaroy/newsgo/rss"
)

func main() {
	feedLines, err := filesystem.ReadFeedFile()
	if err != nil {
		fmt.Println("Error reading feed lines:", err)
		return
	}

	fmt.Println(feedLines)
	for _, feed := range feedLines {
		rss.FetchRss(feed)
	}
}
