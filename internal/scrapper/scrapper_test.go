package scrapper

import (
	"fmt"
	"testing"
)

func TestScrapeWeb(t *testing.T) {
	tests := []struct {
		Url      string
		Expected string
	}{{Url: "https://www.scrapethissite.com/pages", Expected: "Learn Web Scraping | Scrape This Site | A public sandbox for learning web scraping"}}

	for _, test := range tests {
		result := ScrapeWeb(test.Url).Title
		if test.Expected != result {
			t.Logf("Test failed got %v, expected %v", result, test.Expected)
			t.Fail()
		}
	}
}

func TestMultipleScrappers(t *testing.T) {
	scrapeChannel := make(chan string, 1000)
	go ScrapperProducer([]string{"https://www.scrapethissite.com/pages"}, scrapeChannel)
	results := ScrapperConsumer(scrapeChannel)
	fmt.Println("Results from scrapper", results)
}
