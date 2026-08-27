package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Order struct {
	FoodName   string `json:"foodname"`
	TotalCount int    `json:"totalcount"`
}

func orderHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	food := r.URL.Query().Get("food")
	countStr := r.URL.Query().Get("count")
	countNum, err := strconv.Atoi(countStr)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	order := Order{
		FoodName:   food,
		TotalCount: countNum,
	}

	json.NewEncoder(w).Encode(order)
}

func main() {
	http.HandleFunc("/order", orderHandler)

	http.ListenAndServe(":8080", nil)
}
