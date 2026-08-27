package main

import (
	"fmt"
	"net/http"
)

type NewUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprint(w, "Этот маршрут принимает только POST запросы!")
		return
	}
}


var user NewUser

err := json.NewDecoder(r.Body).Decode(&user)
if err != nil {
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprint(w, "Ошибка: Мы не смогли прочитать твой JSON!")
	return
} 	


fmt.Printf("Сервер успешно сохранил пользователя: %s с паролем: %s\n", user.Email, user.Password)


w.Header().Set("Content-Type", "application/json")
w.WriteHeader(http.StatusCreated)


response := map[string]string{"status": "success", "message": "Пользователь зарегестрирован"}
json.NewEncoder(w).Encode(response)

func main() {
	http.HandleFunc("/register", createUserHandler)
	http.ListenAndServe(":8080", nil)
}
