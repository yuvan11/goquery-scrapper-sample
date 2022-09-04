package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func ExampleScrape() {
	// Request the HTML page.
	res, err := http.Get("http://journaldev.com")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	data, _ := doc.Find(".SearchIndexStyles__StyledContainer-sc-msngg5-0").Html()

	fmt.Println(data)
	doc.Find(".SearchIndexStyles__StyledContainer-sc-msngg5-0").Each(func(i int, s *goquery.Selection) {
		heading := s.Find("h3").Text()
		description := s.Find("p").Text()
		fmt.Printf("Review %d: %s ---------- : %s\n", i, heading, description)
	})
}

func main() {
	ExampleScrape()
}
