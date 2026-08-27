package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv" // Обычный стандартный пакет
)

type PriceResponse struct {
	InitialPrice int `json:"initialprice"`
	FinalPrice   int `json:"finalprice"`
}

func discountHandler(w http.ResponseWriter, r *http.Request) {
	// Не забываем скобки () после Header
	w.Header().Set("Content-Type", "application/json")

	// 1. Получаем и переводим Исходную цену (строка в кавычках "price")
	initStrPrice := r.URL.Query().Get("price")
	initIntPrice, err := strconv.Atoi(initStrPrice)
	if err != nil {
		fmt.Println("Ошибка цены:", err)
		return
	}

	// 2. Получаем и переводим Процент скидки (строка в кавычках "percent")
	percentStr := r.URL.Query().Get("percent")
	percentInt, err := strconv.Atoi(percentStr) // Обязательно принимаем err
	if err != nil {
		fmt.Println("Ошибка процента:", err)
		return
	}

	// 3. Считаем финальную цену на языке математики Go
	finalPrice := initIntPrice - (initIntPrice * percentInt / 100)

	// 4. Заполняем структуру
	priceResponse := PriceResponse{
		InitialPrice: initIntPrice,
		FinalPrice:   finalPrice,
	}

	// 5. Отправляем JSON
	json.NewEncoder(w).Encode(priceResponse)
}

func main() {
	http.HandleFunc("/discount", discountHandler)
	http.ListenAndServe(":8080", nil)
}
