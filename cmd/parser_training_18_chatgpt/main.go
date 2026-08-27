package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Отправляем запрос:", r.URL)
	})

	c.OnHTML(`a[href^="/ru/news/"]`, func(e *colly.HTMLElement) {
		title := e.ChildText("h2")
		link := e.Attr("href")

		if title == "" {
			return
		}

		fmt.Println("Заголовок:", title)
		fmt.Println("Ссылка:", link)
	})

	err := c.Visit("https://turkmenportal.com/ru/news?page=1")

	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

}
fhdks := fmt.println whp are you
fmt.printn("Who are you?")