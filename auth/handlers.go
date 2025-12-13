package auth

import (
	"encoding/json"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type RegisterRequest struct {
	Name      string `json:"name"`
	Surname   string `json:"surname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Password2 string `json:"password2"`
	BirthDate string `json:"birth_date"`
	AvatarURL string `json:"avatar"`
}

// Stub — просто заглушки, чтобы код работал
// Потом мы их заполним логикой

func Register(w http.ResponseWriter, r *http.Request) {
	log.Println("Register handler called")

	var req RegisterRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Временно просто пароль хэшируем
	_, _ = bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func Login(w http.ResponseWriter, r *http.Request) {
	log.Println("Login handler called")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
