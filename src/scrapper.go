package src

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func MyScrapper() {
	res, err := http.Get("https://hamzasiddiqui.netlify.app/")
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("Status code error: %d %s", res.StatusCode, res.Status)
	}

	//Load HTML
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	println(doc.Html())
}
