package main

import "github.com/mmcdole/gofeed"

func retrieveFeeds(feedLinks []string) []*gofeed.Feed {
	var feeds []*gofeed.Feed

	for _, link := range feedLinks {
		feed, err := parseFeed(link)
		if err != nil {
			panic(err)
		}
		feeds = append(feeds, feed)
	}
	return feeds
}

func parseFeed(feedLink string) (*gofeed.Feed, error) {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(feedLink)

	return feed, err
}
