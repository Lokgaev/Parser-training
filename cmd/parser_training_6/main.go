package main

import (
	"encoding/json"
	"net/http"
)

type Developer struct {
	Name     string `json:"name"`
	Language string `json:"language"`
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	dev := Developer{
		Name:     "Muhammed",
		Language: "English",
	}

	json.NewEncoder(w).Encode(dev)
}

func main() {
	http.HandleFunc("/", mainHandler)

	http.ListenAndServe(":8080", nil)
}
