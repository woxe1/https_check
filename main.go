package main

import (
	"log"
	"net/http"
)

func mainHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, HTTPS with PEM!"))
}

func main() {
	http.HandleFunc("/", mainHandler)

	// Запуск HTTPS сервера на порту 443 с использованием SSL сертификата и закрытого ключа в формате PEM
	err := http.ListenAndServeTLS(":443", "fullchain.pem", "privkey.pem", nil)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
