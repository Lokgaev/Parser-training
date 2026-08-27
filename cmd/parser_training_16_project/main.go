package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

type FullArticle struct {
	Title    string `json:"title"`
	URL      string `json:"url"`
	FullText string `json:"full_text"`
}

func main() {
	// 1. Главный коллектор списков
	c := colly.NewCollector(
		colly.AllowedDomains("turkmenportal.com"),
	)
	c.Limit(&colly.LimitRule{
		DomainRegexp: `turkmenportal\.com`,
		Delay:        1 * time.Second, // Задержка 1 сек, чтобы не перегружать сайт
	})

	// 2. Второй коллектор для чтения текста внутри статьи
	articleCollector := c.Clone()

	// ВНУТРИ СТАТЬИ: забираем весь текст
	articleCollector.OnHTML("main", func(e *colly.HTMLElement) {
		title := e.ChildText("h1")

		var paragraphs []string
		e.ForEach("p", func(_ int, p *colly.HTMLElement) {
			text := strings.TrimSpace(p.Text)
			if text != "" {
				paragraphs = append(paragraphs, text)
			}
		})

		article := FullArticle{
			Title:    title,
			URL:      e.Request.URL.String(),
			FullText: strings.Join(paragraphs, "\n\n"),
		}

		fmt.Printf("\n=== ПОЛНАЯ СТАТЬЯ СПАРСЕНА ===\nЗаголовок: %s\nURL: %s\nДлина текста: %d символов\n",
			article.Title, article.URL, len(article.FullText))
	})

	// НА СТРАНИЦЕ СПИСКА: находим карточки и отправляем articleCollector
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		title := e.ChildText("h2")

		if title != "" {
			fullURL := e.Request.AbsoluteURL(e.Attr("href"))

			// ЗАПУСКАЕМ скачивание полной статьи по найденной ссылке!
			articleCollector.Visit(fullURL)
		}
	})

	fmt.Println("Начинаем парсинг...")
	c.Visit("https://turkmenportal.com/ru/news?page=1")
}
