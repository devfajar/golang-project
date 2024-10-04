package main

import (
	"bytes"
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

func main() {
	url := "https://mangaread.org"

	c := colly.NewCollector()
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting...", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		reader := bytes.NewReader(r.Body)
		doc, err := goquery.NewDocumentFromReader(reader)
		if err != nil {
			log.Fatal("Error loading HTTP response body: ", err)
		}
		title := doc.Find("Title").Text()
		fmt.Println("Title:", title)
	})

	// Visit the URL
	err := c.Visit(url)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
