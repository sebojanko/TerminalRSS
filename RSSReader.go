package main

import (
	"github.com/mmcdole/gofeed"
	"log"
)

var Cleaners map[string]string

func init() {
	Cleaners = getCleaners()
}

func main() {
	var feeds []*gofeed.Feed

	//TODO feed se ucitava tek kad otvorimo taj screen

	feedLinks := getLinks()
	feeds = retrieveFeeds(feedLinks)

	ui := createTUI(feeds)

	if err := ui.Run(); err != nil {
		log.Fatal(err)
	}
}
