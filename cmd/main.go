package main

import (
	"flag"
	"fmt"
	"strings"
	"webscrapper/internal/scrapper"
)

func main() {
	var urlsInput string
	flag.StringVar(&urlsInput, "urls", "https://example.com", "URL to scrape")

	flag.Parse()

	urlsToScrape := strings.Split(urlsInput, ",")

	scrapeChannel := make(chan string, 1000)
	go scrapper.ScrapperProducer(urlsToScrape, scrapeChannel)
	results := scrapper.ScrapperConsumer(scrapeChannel)

	for _, result := range results {
		fmt.Println(result)
	}
}
