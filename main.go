package main

import (
	"log"
	"os"

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

 	c := colly.NewCollector()

 	// Define the URL you want to scrape
 	url := "https://www.deine-berge.de/POIs/Filter/Kategorie-1-Berg-Gipfel+Gebirge-13-Chiemgauer-Alpen/Alle"

	println("I am here")

	c.OnHTML("html", func(e *colly.HTMLElement) {
    	println("Found an HTML tag!")
	})

	// Visit the URL and start scraping
	err = c.Visit(url)
	if err != nil {
		log.Fatal(err)
	}
}