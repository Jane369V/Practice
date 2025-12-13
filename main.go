package main

import (
	"log"
	"net/http"

	"practice/auth"
)

func main() {
	auth.InitDB()
	log.Println("База данных подключена")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server is alive 🚀"))
	})

	http.HandleFunc("/register", auth.Register)
	http.HandleFunc("/login", auth.Login)

	log.Println("Сервер запущен на :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
