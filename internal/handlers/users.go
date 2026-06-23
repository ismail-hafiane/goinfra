package handlers

import (
	"encoding/json"
	"net/http"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status": "ok",
		"server": "goinfra",
	})
}

func Users(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		users := []User{
			{ID: 1, Name: "Ismail", Email: "ismail@goinfra.com"},
			{ID: 2, Name: "Admin", Email: "admin@goinfra.com"},
		}
		json.NewEncoder(w).Encode(users)

	case http.MethodPost:
		var user User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
		user.ID = 3
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(user)

	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
