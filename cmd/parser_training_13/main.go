package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Item struct {
	Name  string `json:"name"`
	Stock int    `json:"stock"`
}

type LoginData struct {
	Login string `json:"login"`
	Pass  string `json:"pass"`
}

func addProductHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprint(w, "Ошибка: метод не является POST")
		return
	}

	var item Item

	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Ошибка чтения JSON: %v", err)
		return
	}

	fmt.Printf("Товар успешно добавлен на склад! Имя: %s, Количество: %d\n", item.Name, item.Stock)

	w.Header().Set("Content-Type", "application/json")

	response := map[string]string{"status": "added"}
	json.NewEncoder(w).Encode(response)

}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprint(w, "Ошибка: метод не является POST")
		return
	}

	var loginData LoginData

	err := json.NewDecoder(r.Body).Decode(&loginData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Ошибка чтения JSON: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	response := make(map[string]string)

	if loginData.Login == "admin" && loginData.Pass == "1234" {
		response["auth"] = "success"
		w.WriteHeader(http.StatusOK)
	} else {
		response["auth"] = "denied"
		w.WriteHeader(http.StatusUnauthorized)
	}

	json.NewEncoder(w).Encode(response)

}

func main() {
	http.HandleFunc("/api/add-product", addProductHandler)
	http.HandleFunc("/api/login", loginHandler)
	fmt.Println("Сервер склада запущен на порту 808...")
	http.ListenAndServe(":8080", nil)
}
