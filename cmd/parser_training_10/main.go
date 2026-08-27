package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Player struct {
	User  string `json:"user"`
	Level int    `json:"level"`
}

func profileHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	user := r.URL.Query().Get("nickname")
	levelStr := r.URL.Query().Get("lvl")
	levelInt, err := strconv.Atoi(levelStr)
	if err != nil {
		fmt.Println("Ошибка:", err)
	}
	player := Player{
		User:  user,
		Level: levelInt,
	}

	json.NewEncoder(w).Encode(player)

}

func main() {
	http.HandleFunc("/profile", profileHandler)

	http.ListenAndServe(":8080", nil)

}
