package src

import (
	"net/http"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
)

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

		results <- Result{URL: url, StatusCode: res.StatusCode}

		println(doc.Html())

		time.Sleep(1 * time.Second)
	}
}
