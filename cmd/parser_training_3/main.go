package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

func main() {
	// Создаём паука-коллектора
	c := colly.NewCollector()

	// Говорим пауку: "Когда найдёшь блок с цитатой, сделай следующее:"
	c.OnHTML("div.quote", func(e *colly.HTMLElement) {

		// 1. Достань автора (текст из small.author) и положи в переменную author
		author := e.ChildText("small.author")

		// 2. Создай пустой список для тегов
		tags := []string{}

		// 3. Пройдись по всем a.tag внутри этой цитаты
		e.ForEach("a.tag", func(i int, elem *colly.HTMLElement) {
			// Для каждого найденного тега: возьми его текст и добавь в список
			tags = append(tags, elem.Text)
		})

		// 4. Преврати список тегов в строку через запятую
		tagsString := strings.Join(tags, ", ")

		// 5. Напечатай результат (теперь с аргументами после запятой!)
		fmt.Printf("Автор: %s\n", author)
		fmt.Printf("Теги: %s\n", tagsString)
		fmt.Println("---")
	})

	// Запускаем паука на 3 страницы подряд
	for page := 1; page <= 3; page++ {
		// Составляем URL, подставляя номер страницы
		url := fmt.Sprintf("http://quotes.toscrape.com/page/%d/", page)

		// Сообщаем, какую страницу парсим
		fmt.Println("=== Парсим страницу:", url, "===")

		// Отправляем паука на эту страницу
		c.Visit(url)
	}
}
