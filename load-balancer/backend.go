// backend.go
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func startBackend(port string) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Ответ от бэкенда на порту %s\n", port)
	})

	log.Printf("Запуск бэкенда на порту %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// Функция main() для запуска сервера
func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "9001" // По умолчанию если переменная PORT не установлена
	}

	startBackend(port)
}
