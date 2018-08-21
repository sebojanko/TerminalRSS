package main

import (
	"encoding/json"
	"io/ioutil"
)

const CONFIG_PATH string = "links.json"

var links map[string][]string

func readLinksFile() []byte {
	dat, err := ioutil.ReadFile(CONFIG_PATH)
	if err != nil {
		panic(err)
	}

	return dat
}

func getLinks() []string {
	data := readLinksFile()

	json.Unmarshal(data, &links)
	return links["links"]
}
