package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gocolly/colly/v2"
)


func main() {
	fName := "heights.csv"
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Cannot create file %q: %s\n", fName, err)
		return
	}
	defer file.Close()

	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: coursera.org, www.coursera.org
		colly.AllowedDomains("https://www.deine-berge.de"),

		// Cache responses to prevent multiple download of pages
		// even if the collector is restarted
		colly.CacheDir("./height_cache"),
		// Cached responses older than the specified duration will be refreshed
		colly.CacheExpiration(24*time.Hour),
	)

	fmt.Println("I am here")

	// On every <a> element which has "href" attribute call callback
	c.OnHTML("tr", func(e *colly.HTMLElement) {
		// If attribute class is this long string return from callback
		// As this a is irrelevant
		fmt.Println(e)
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		log.Println("visiting", r.URL.String())
	})

	//summits := make([]Summit, 0, 200)

		// Start scraping on http://coursera.com/browse
	c.Visit("https://www.deine-berge.de/POIs/Filter/Kategorie-1-Berg-Gipfel+Gebirge-13-Chiemgauer-Alpen/Alle")
}