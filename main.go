package main

import (
	"fmt"
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

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Scraping:", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
       	fmt.Println("Status:", r.StatusCode)
   	})

	c.OnHTML("table", func(e *colly.HTMLElement) {
    	fmt.Println("Found table")
		e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			s := Summit{ Name: el.ChildText("td:nth-child(2)") }
			fmt.Printf("%+v\n", s)
		})
	})

	c.OnError(func(r *colly.Response, err error) {
    	fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
   	})

	// Visit the URL and start scraping
	err = c.Visit(url)
	if err != nil {
		log.Fatal(err)
	}
}