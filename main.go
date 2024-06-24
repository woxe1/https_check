package main

import (
	"log"
	"net/http"
)

// CORS Middleware
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight request
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, HTTPS with PEM!"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", mainHandler)

	// Оборачиваем основной маршрутизатор в CORS middleware
	handler := corsMiddleware(mux)

	// Запуск HTTPS сервера на порту 443 с использованием SSL сертификата и закрытого ключа в формате PEM
	err := http.ListenAndServeTLS(":443", "fullchain.pem", "privkey.pem", handler)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
