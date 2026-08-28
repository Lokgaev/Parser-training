package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

type News struct {
	Title string
	URL   string
	Text  string
}

func main() {
	listCollector := colly.NewCollector()
	articleCollector := colly.NewCollector()

	var newsList []News

	var currentTitle string
	var currentURL string
	var currentText strings.Builder

	listCollector.OnHTML(`a[href^="/ru/news/"]`, func(e *colly.HTMLElement) {
		title := e.ChildText("h2")

		if title == "" {
			return
		}

		link := e.Attr("href")
		fullURL := e.Request.AbsoluteURL(link)

		currentTitle = title
		currentURL = fullURL

		currentText.Reset()

		err := articleCollector.Visit(fullURL)

		if err != nil {
			fmt.Println("Ошибка статьи:", err)
		}
	})

	articleCollector.OnHTML("#content p", func(e *colly.HTMLElement) {
		text := strings.TrimSpace(e.Text)

		if text == "" {
			return
		}

		currentText.WriteString(text)
		currentText.WriteString("\n")
	})

	articleCollector.OnScraped(func(r *colly.Response) {
		news := News{
			Title: currentTitle,
			URL:   currentURL,
			Text:  strings.TrimSpace(currentText.String()),
		}

		newsList = append(newsList, news)
	})

	err := listCollector.Visit("https://turkmenportal.com/ru/news?page=1")

	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	fmt.Println("Количество новостей:", len(newsList))

	for _, news := range newsList {
		fmt.Println("Заголовок:", news.Title)
		fmt.Println("URL:", news.URL)
		fmt.Println("Текст:", news.Text)
		fmt.Println("----------------")
	}
}
