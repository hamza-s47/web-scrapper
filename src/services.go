package src

import (
	"log"
	"sync"
)

func ScrapeService(urls Urls) []Result {

	jobs := make(chan string, 5)
	results := make(chan Result, 5)
	var finalResult []Result

	var wg sync.WaitGroup

	// Workers Starting (3 workers)
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go Worker(i, jobs, results, &wg)
	}

	// Producer
	go func() {
		for _, url := range urls.URLs {
			jobs <- url
		}
		close(jobs)
	}()

	// Closing result channel
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collector
	for res := range results {
		if res.Err != nil {
			log.Fatal(res.Err)
			continue
		}
		finalResult = append(finalResult, res)
	}
	return finalResult
}
