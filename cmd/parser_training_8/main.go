package main

import (
	"encoding/json"
	"net/http"
)

type Game struct {
	Title string `json:"title"`
	Genre string `json:"genre"`
}

func gamesHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	games := []Game{
		{Title: "Dishonored", Genre: "stealth-action"},
		{Title: "Fear", Genre: "action-horror"},
	}

	json.NewEncoder(w).Encode(games)
}

func main() {
	http.HandleFunc("/games", gamesHandler)

	http.ListenAndServe(":8080", nil)

}
