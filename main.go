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
		fmt.Printf("%s%sFor helping instruction type:\n", src.Bold, src.Red)
		// fmt.Println("./app http://url1.com http://url2.com ... http://url.com")
		fmt.Printf("./app --help%s\n\n\n", src.Reset)
	}

	switch args[0] {
	case "--api":
		r := gin.Default()
		src.RegisterRoutes(r)

	case "--help", "-h":
		fmt.Printf("%s%sType the following commands (green ones):%s\n", src.BU, src.Blue, src.Reset)
		fmt.Printf("%s%s./app http://url1.com http://url2.com ... http://url.com %s----> %s(Run scraper via CLI)%s\n", src.Green, src.Bold, src.Blue, src.Yellow, src.Reset)
		fmt.Printf("%s%s./app --api	%s-------------------------------------------->  %s(Start API server on port: 8080)%s\n", src.Green, src.Bold, src.Blue, src.Yellow, src.Reset)

	default:
		src.ScrapeCli(args)
	}

}
