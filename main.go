package main

import (
	"crypto/tls"
	"log"
	"net/http"
)

func main() {
	// Загрузка сертификата и ключа
	cert, err := tls.LoadX509KeyPair("./certificate.crt", "./private.key")
	if err != nil {
		log.Fatalf("failed to load key pair: %s", err)
	}

	// Настройка TLS конфигурации
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}

	// Создание HTTPS сервера
	server := &http.Server{
		Addr:      ":443",
		TLSConfig: tlsConfig,
	}

	// Определение обработчика для корневого пути
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, HTTPS!"))
	})

	log.Println("Starting HTTPS server on :443")
	err = server.ListenAndServeTLS("", "")
	if err != nil {
		log.Fatalf("server failed to start: %s", err)
	}
}
