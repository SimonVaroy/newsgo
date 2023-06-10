package main

import (
	"github.com/SimonVaroy/newsgo/rss"
)

func main() {
	rss.FetchRss("https://lukesmith.xyz/index.xml")
}
