package src

import (
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func extractLinks(doc *goquery.Document) ([]string, int) {
	links := []string{}

	doc.Find("a[href]").Each(func(i int, s *goquery.Selection) {
		href, exist := s.Attr("href")
		if exist && strings.TrimSpace(href) != "" {
			links = append(links, strings.TrimSpace(href))
		}
	})

	return links, len(links)
}

func extractHeadings(doc *goquery.Document) ([]string, int) {
	headings := []string{}

	doc.Find("h1, h2, h3, h4, h5, h6").Each(func(i int, s *goquery.Selection) {
		text := strings.TrimSpace(s.Text())
		if text != "" {
			headings = append(headings, text)
		}
	})

	return headings, len(headings)
}

func Worker(id int, jobs <-chan string, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()

	client := &http.Client{Timeout: 10 * time.Second}

	for url := range jobs {

		res, err := client.Get(url)
		if err != nil {
			results <- Result{URL: url, Err: err}
			continue
		}

		//Load HTML
		doc, err := goquery.NewDocumentFromReader(res.Body)
		res.Body.Close()

		if err != nil {
			results <- Result{URL: url, Err: err}
			continue
		}

		links, lc := extractLinks(doc)
		headings, hc := extractHeadings(doc)

		data := ScrapedData{Links: links, Headings: headings}

		results <- Result{URL: url, StatusCode: res.StatusCode, LinksCount: lc, HeadingsCount: hc, Data: data}

		time.Sleep(1 * time.Second)
	}
}
