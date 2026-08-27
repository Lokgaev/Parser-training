package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gocolly/colly"
)


type Article struct {
	Title string
	FullText string
	Tags []string
}


func main() {
	c := colly.NewCollector()

	absPath, _ := filepath.Abs("index.html")

	c.OnHTML(".news-article", func(e *colly.HTMLElement){
		e.ChildText(".article-title")

	})

	c.OnHTML(".article-body" func(e *colly.HTMLElement){
		e.ChildText("p")
	})


	c.Visit("turkmenportal.com/news/ru")
}