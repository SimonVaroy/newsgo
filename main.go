package main

import (
	"fmt"

	"github.com/SimonVaroy/newsgo/rss"
)

func main() {
	feedLines, err := rss.ReadFeedFile()
	if err != nil {
		fmt.Println("Error reading feed lines:", err)
		return
	}

	fmt.Println(feedLines)
}
