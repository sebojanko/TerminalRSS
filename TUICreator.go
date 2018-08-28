package main

import (
	"github.com/marcusolsson/tui-go"
	"github.com/mmcdole/gofeed"
	"log"
	"os/exec"
	"regexp"
	"strconv"
)

func createTUI(feeds []*gofeed.Feed) tui.UI {
	var currView int
	var tables []*tui.Table
	var labels []*tui.Label
	var boxes []*tui.Box
	var views []tui.Widget

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

	return ui
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
		articleDescription = cleanup(feed.Title, articleDescription)
		articleDescription = htmlTags.ReplaceAllString(articleDescription, "")

		label.SetText(articleDescription)
	})
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
