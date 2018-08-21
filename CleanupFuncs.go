package main

import "regexp"

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
