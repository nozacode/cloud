package main

import (
	"load-balancer/balancer"
	"load-balancer/handler"
	"log"
	"net/http"
	"time"
)

func main() {
	// Инициализируем пул бэкендов
	backends := []string{
		"http://localhost:8081",
		"http://localhost:8082",
		"http://localhost:8083",
	}

	// Создаём экземпляр балансировщика
	lb := balancer.NewLoadBalancer(backends)

	// Запускаем периодическую проверку доступности бэкендов
	go lb.HealthCheck(10 * time.Second)

	// Назначаем обработчик для всех входящих путей
	http.HandleFunc("/", handler.NewProxyHandler(lb))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
	log.Println("Load Balancer запущен на порту :8080")
}

/*package main

import (
	"load-balancer/handler"
	"log"
	"net/http"
	"time"
)

func main() {
	// Инициализируем балансировщик с пулом бэкендов
	backends := []string{
		"http://localhost:8081",
		"http://localhost:8082",
		"http://localhost:8083",
	}
	// Инициализируем балансировщик
	go lb.HealthCheck(10 * time.Second)

	// Назначаем обработчик для всех входящих путей
	http.HandleFunc("/", handler.NewProxyHandler(lb))

	log.Println("Load Balancer запущен на порту :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}*/
