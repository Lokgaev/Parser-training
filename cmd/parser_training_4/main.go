package main

import (
	"fmt"
	"net/http"
)

func main() {
	response, err := http.Get("https://example.com")
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	fmt.Printf("Статус: %s\n", response.Status)
	fmt.Printf("%d\n", response.StatusCode)
	response.Body.Close()
}
