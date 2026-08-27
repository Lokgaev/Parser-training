package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("https://example.com")
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	fmt.Printf("Статус: %s\n", resp.Status)
	fmt.Printf("Код статуса: %d\n", resp.StatusCode)
	fmt.Println("Все headers:")
	fmt.Println(resp.Header)

	fmt.Println("Content-Type:")
	fmt.Println(resp.Header.Get("Content-Type"))

	// Перебираем все HTTP-заголовки.
	//
	// resp.Header имеет тип http.Header,
	// а http.Header основан на:
	//
	// map[string][]string
	//
	// Поэтому key — название заголовка,
	// values — его значения.

	for key, values := range resp.Header {
		fmt.Printf("%s, %v\n", key, values)
	}
	resp.Body.Close()
}
