package scrapper

import (
	"log"
	"sync"

	"github.com/gocolly/colly/v2"
)

type ScrapedData struct {
	URL   string
	Title string
}

const (
	Title = "title"
)

func ScrapeWeb(url string) ScrapedData {
	c := colly.NewCollector(
		colly.AllowedDomains("www.scrapethissite.com"),
		colly.IgnoreRobotsTxt(),
	)

	var title string

	c.OnHTML(Title, func(h *colly.HTMLElement) {
		title = h.Text
	})

	// Set a browser-like User-Agent
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36"

	err := c.Visit(url)

	if err != nil {
		log.Printf("Error visiting %s: %v", url, err)
		return ScrapedData{}
	}

	return ScrapedData{
		URL:   url,
		Title: title,
	}
}

func ScrapperProducer(urls []string, scrapperChannel chan string) {
	for _, url := range urls {
		scrapperChannel <- url
	}
	close(scrapperChannel)
}

func ScrapperConsumer(scrapperChannel chan string) []string {
	var wg sync.WaitGroup
	results := make([]string, 0)
	mutex := sync.Mutex{}

	for url := range scrapperChannel {
		wg.Add(1) // Increment for each goroutine
		go func(url string) {
			defer wg.Done() // Decrement when goroutine finishes
			result := ScrapeWeb(url)
			mutex.Lock()
			results = append(results, result.Title)
			mutex.Unlock()
		}(url)
	}

	wg.Wait() // Wait for all goroutines to finish
	return results
}
