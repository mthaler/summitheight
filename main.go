package main

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
}