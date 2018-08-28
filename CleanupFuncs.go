package main

import (
	"regexp"
)

func cleanup(title, articleDescription string) string {
	cleanupPattern, err := regexp.Compile(Cleaners[title])
	if err != nil {
		panic(err)
	}
	
	articleDescription = cleanupPattern.ReplaceAllString(articleDescription, "")
	return articleDescription
}
