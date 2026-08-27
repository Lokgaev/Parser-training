package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

type LibraryResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func addBookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintln(w, "Ошибка: метод не является POST")
		return
	}

	var book Book

	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Ошибка чтения JSON: %v\n", err)
		return
	}

	fmt.Printf("Книга: %s от автора: %s успешно добавлена!", book.Title, book.Author)

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)

	response := LibraryResponse{
		Status:  "success",
		Message: "Книга добавлена",
	}

	json.NewEncoder(w).Encode(response)

}

func searchBookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintln(w, "Ошибка: метод не является GET")
		return
	}

	titleQuery := r.URL.Query().Get("title")
	if titleQuery == "Hobbit" {
		w.Header().Set("Content-Type", "application/json")
		book := Book{Title: "Hobbit", Author: "Tolkien"}
		json.NewEncoder(w).Encode(book)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		libResponse := LibraryResponse{Status: "error", Message: "Книга не найдена"}
		json.NewEncoder(w).Encode(libResponse)
	}

}

func main() {
	http.HandleFunc("/api/add-book", addBookHandler)
	http.HandleFunc("/api/search-book", searchBookHandler)
	http.ListenAndServe(":8080", nil)
}
