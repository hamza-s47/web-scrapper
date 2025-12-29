package src

import (
	"fmt"
	"strconv"
)

const (
	Red       = "\033[31m"
	Green     = "\033[32m"
	Yellow    = "\033[33m"
	Blue      = "\033[34m"
	Bold      = "\033[1m"
	Underline = "\033[4m"
	BU        = "\033[1m \033[4m"
	Reset     = "\033[0m"
)

func ScrapeCli(args []string) {

	cliArgs := Urls{
		URLs: args,
	}

	results := ScrapeService(cliArgs)

	for _, r := range results {
		fmt.Println(Blue+BU+"URL:"+Reset+Bold, Green+r.URL+Reset)
		fmt.Println(Blue+BU+"Status:"+Reset+Bold, Green+strconv.Itoa(r.StatusCode)+Reset)
		if r.Err != nil {
			fmt.Println(Yellow+BU+"Error:"+Reset+Bold+Red, r.Err)
		}
		fmt.Println(Blue+BU+"Links:"+Reset+Bold+Green, len(r.Data.Links))
		fmt.Println(Blue+BU+"Headings:"+Reset+Bold+Green, len(r.Data.Headings))
		fmt.Println(Reset + Red + Bold + "-----------------------------------------" + Reset)
		fmt.Println(Blue+BU+"Links:"+Reset+Bold+Green, r.Data.Links)
		fmt.Println("")
		fmt.Println(Blue+BU+"Headings:"+Reset+Bold+Green, r.Data.Headings)
	}
}
