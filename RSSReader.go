package main

import (
	"github.com/mmcdole/gofeed"
	"github.com/marcusolsson/tui-go"
	"log"
	"strconv"
	"regexp"
	"os/exec"
)

func main() {
	var currView int
	var feeds []*gofeed.Feed
	var tables []*tui.Table
	var labels []*tui.Label
	var boxes []*tui.Box
	var views []tui.Widget

	//TODO feed se ucitava tek kad otvorimo taj screen
	//TODO feed linkovi se nalaze u fileu
	feedLinks :=
		[]string{
			"https://www.vecernji.hr/feeds/latest",
			"https://www.muzika.hr/feed/",
			"https://www.muzika.hr/feed/?cat=11",
			"https://www.muzika.hr/feed/?cat=18",
			"https://www.muzika.hr/feed/?cat=3",
			"https://www.muzika.hr/feed/?cat=5",
		}

	for _, link := range feedLinks {
		feed, err := parseFeed(link)
		if err != nil {
			panic(err)
		}
		feeds = append(feeds, feed)
	}

	for _, feed := range feeds {
		table := createTable()
		table = fillTable(table, feed)
		tables = append(tables, table)

		label := createLabel()
		labels = append(labels, label)

		box := createScreen(table, label, feed.Title)
		boxes = append(boxes, box)

		views = append(views, box)

		onSelectionChanged(table, feed, label)
		onItemActivated(table, feed)
	}

	ui, err := tui.New(views[0])
	if err != nil {
		log.Fatal(err)
	}

	ui.SetKeybinding("Esc", func() { ui.Quit() })
	ui.SetKeybinding("q", func() { ui.Quit() })
	ui.SetKeybinding("Left", func() {
		currView = clamp(currView-1, 0, len(views)-1)
		ui.SetWidget(views[currView])
	})
	ui.SetKeybinding("Right", func() {
		currView = clamp(currView+1, 0, len(views)-1)
		ui.SetWidget(views[currView])
	})

	if err := ui.Run(); err != nil {
		log.Fatal(err)
	}
}
func onItemActivated(table *tui.Table, feed *gofeed.Feed) {
	table.OnItemActivated(func(t *tui.Table) {
		item := t.Selected()
		articleLink := feed.Items[item].Link
		cmd := exec.Command("firefox", articleLink)
		cmd.Start()
	})
}
func onSelectionChanged(table *tui.Table, feed *gofeed.Feed, label *tui.Label) {
	table.OnSelectionChanged(func(t *tui.Table) {
		item := t.Selected()
		htmlTags, _ := regexp.Compile("<.*?>")
		articleDescription := feed.Items[item].Description

		articleDescription = additionalCleanup(feed.Title, articleDescription)

		articleDescription = htmlTags.ReplaceAllString(articleDescription, "")
		label.SetText(articleDescription)
	})
}
func additionalCleanup(title string, articleDescription string) string {
	if title == "muzika.hr" {
		articleDescription = cleanUpMuzikaHr(articleDescription)
	}
	return articleDescription
}
func cleanUpMuzikaHr(articleDescription string) string {
	objavaTags, _ := regexp.Compile("Objava <.*?>.*</p>")
	articleDescription = objavaTags.ReplaceAllString(articleDescription, "")
	return articleDescription
}
func createScreen(t *tui.Table, l *tui.Label, title string) *tui.Box {
	tableBox := tui.NewVBox(t)
	tableBox.SetBorder(true)
	labelBox := tui.NewVBox(l)
	labelBox.SetBorder(true)
	rootBox := tui.NewVBox(tableBox, labelBox)
	tableBox.SetTitle(title)
	return rootBox
}
func createLabel() *tui.Label {
	article := tui.NewLabel("")
	article.SetWordWrap(true)
	return article
}

func createTable() *tui.Table {
	table := tui.NewTable(0, 0)
	table.SetFocused(true)
	return table
}

func fillTable(t *tui.Table, items *gofeed.Feed) *tui.Table {

	for i, item := range items.Items {
		t.AppendRow(tui.NewLabel(strconv.Itoa(i) + " " + item.Title))
	}

	return t
}
func parseFeed(feedLink string) (*gofeed.Feed, error) {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(feedLink)

	return feed, err
}
func clamp(n, min, max int) int {
	if n < min {
		return max
	}
	if n > max {
		return min
	}
	return n
}
