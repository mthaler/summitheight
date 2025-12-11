package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

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

	var summits []Summit
	summits = make([]Summit, 1000)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Scraping:", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
       	fmt.Println("Status:", r.StatusCode)
   	})

	c.OnHTML("table", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			heightText := el.ChildText("td:nth-child(3)")
			if (heightText != "") {
				heightText = heightText[:len(heightText)-2]
			}
			height, err := strconv.Atoi(heightText)
			if err != nil {
				fmt.Printf("Could not convert %s to int")
			}
			s := Summit{ Name: el.ChildText("td:nth-child(2)"),
				 Kategorie: el.ChildText("td:nth-child(4)"),
				 Height: height,
				 Staat: el.ChildText("td:nth-child(5)"),
				 Region: el.ChildText("td:nth-child(6)"),
				 Gruppe: el.ChildText("td:nth-child(7)"),
				 Information: el.ChildText("td:nth-child(8)"),
			}
			summits = append(summits, s)
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

	w := csv.NewWriter(file)
	defer w.Flush()

	w.WriteAll(summits)
}