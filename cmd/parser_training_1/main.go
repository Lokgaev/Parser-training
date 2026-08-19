package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()

	// Ищем общий блок div с классом quote
	c.OnHTML("div.quote", func(e *colly.HTMLElement) {
		// Внутри этого блока (e) мы ищем ребенка span с классом text
		quote := e.ChildText("span.text")

		// Внутри этого же блока (e) мы ищем ребенка small с классом author
		author := e.ChildText("small.author")

		// Теперь у нас есть и то, и другое, и мы можем их вывести вместе!
		fmt.Printf("%s: \"%s\"\n", author, quote)
	})

	c.Visit("http://quotes.toscrape.com/")
}
