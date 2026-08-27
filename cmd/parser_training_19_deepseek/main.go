package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Собираемся посетить:", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		if r.StatusCode != 200 {
			fmt.Println("Ошибка, статус:", r.StatusCode)
			r.Request.Abort()
		}
	})

	c.OnHTML("a", func(e *colly.HTMLElement) {
		text := e.Text
		href := e.Attr("href")
		fmt.Printf("Текст: %s, ссылка: %s\n", text, href)
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Готово!")
	})

	c.Visit("https://example.com")
}
