package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()

	c.OnHTML("span.text", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
	})

	c.Visit("https://quotes.toscrape.com/")

}
