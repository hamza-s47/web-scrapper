package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hamza-s47/web-scrapper/src"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Type:")
		fmt.Println("./app http://url1.com http://url2.com ... http://url.com")
		fmt.Println("./app main.go --help")
	}

	switch args[0] {
	case "--api":
		r := gin.Default()
		src.RegisterRoutes(r)

	case "--help", "-h":
		fmt.Println("Type:")
		fmt.Println("./app http://url1.com http://url2.com ... http://url.com		(Run scraper via CLI)")
		fmt.Println("./app --api    (Start API server on port: 8080)")

	default:
		src.ScrapeCli(args)
	}

}
