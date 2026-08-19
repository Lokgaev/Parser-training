package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly" // Важно: v2!
)

func main() {
	c := colly.NewCollector()

	c.OnHTML("div.quote", func(e *colly.HTMLElement) {

		// 1. Автор
		author := e.ChildText("small.author")

		// 2. Ссылка
		relativeLink := e.ChildAttr("a", "href")
		fullLink := e.Request.AbsoluteURL(relativeLink)

		// 3. Теги — используем ForEach
		var tags []string // Создаём пустой список строк
		e.ForEach("a.tag", func(i int, elem *colly.HTMLElement) {
			// Эта функция вызовется для каждого найденного a.tag
			// elem.Text — текст текущего тега
			tags = append(tags, elem.Text) // Добавляем текст в список
		})

		tagsString := strings.Join(tags, ", ")

		// 4. Вывод
		fmt.Printf("Автор: %s\n", author)
		fmt.Printf("Ссылка: %s\n", fullLink)
		fmt.Printf("Теги: %s\n", tagsString)
		fmt.Println("----------")
	})

	c.Visit("http://quotes.toscrape.com/")
}
