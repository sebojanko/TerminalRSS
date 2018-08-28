package main

import (
	"encoding/json"
	"io/ioutil"
)

const CLEANERS_PATH string = "cleaners.json"

var cleaners map[string]string

func readCleanersFile() []byte {
	dat, err := ioutil.ReadFile(CLEANERS_PATH)
	if err != nil {
		panic(err)
	}

	return dat
}

func getCleaners() map[string]string {
	data := readCleanersFile()

	json.Unmarshal(data, &cleaners)
	return cleaners
}
