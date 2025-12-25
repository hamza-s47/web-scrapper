package main

import (
	// "github.com/gin-gonic/gin"
	"fmt"
	"log"
	"sync"

	"github.com/hamza-s47/web-scrapper/src"
)

func main() {
	// r := gin.Default()
	// src.RegisterRoutes(r)

	urls := []string{
		"https://github.com/hamza-s47",
		"https://hamzasiddiqui.netlify.app/",
		"https://www.hackerrank.com/profile/hamza_24",
	}

	jobs := make(chan string, 5)
	results := make(chan src.Result, 5)

	var wg sync.WaitGroup

	// Workers Starting
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go src.Worker(i, jobs, results, &wg)
	}

	// Producer
	go func() {
		for _, url := range urls {
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
		fmt.Println("Completed:", res.URL, "Status:", res.StatusCode)
	}
}
