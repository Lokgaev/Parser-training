package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Square struct {
	Side     int `json:"side"`
	Perimetr int `json:"perimetr"`
}

func squareHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	val := r.URL.Query().Get("side")

	sideNum, err := strconv.Atoi(val)
	if err != nil {
		// Если пользователь ввел вместо числа буквы (например, ?side=привет),
		// то strconv.Atoi вернет ошибку. 
		// Давай для безопасности дадим стороне значение по умолчанию, например 1
		sideNum = 1
	}

	square := Square{
		Side:     sideNum,
		Perimetr: 0,
	}

	square.Perimetr = square.Side * sideNum

	json.NewEncoder(w).Encode(square)

}

func main() {
	http.HandleFunc("/square", squareHandler)

	http.ListenAndServe(":8080", nil)
}
